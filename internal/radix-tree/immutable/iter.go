package immutable

import (
	"bytes"
)

// Iterator is used to iterate over a set of nodes
// in pre-order
type Iterator struct {
	node  *Node
	stack []edges
}

// SeekPrefixWatch is used to seek the iterator to a given prefix
// and returns the watch channel of the finest granularity
func (i *Iterator) SeekPrefixWatch(prefix []byte) (watch <-chan struct{}) {
	// Wipe the stack
	i.stack = nil
	n := i.node
	watch = n.mutateCh
	search := prefix
	for {
		// Check for key exhaution
		if len(search) == 0 {
			i.node = n
			return
		}
		// Look for an edge
		_, n = n.getEdge(search[0])
		if n == nil {
			i.node = nil
			return
		}
		// Update to the finest granularity as the search makes progress
		watch = n.mutateCh
		// Consume the search prefix
		if bytes.HasPrefix(search, n.prefix) {
			search = search[len(n.prefix):]
		} else if bytes.HasPrefix(n.prefix, search) {
			i.node = n
			return
		} else {
			i.node = nil
			return
		}
	}
}

// SeekPrefix is used to seek the iterator to a given prefix
func (i *Iterator) SeekPrefix(prefix []byte) {
	i.SeekPrefixWatch(prefix)
}
func (i *Iterator) recurseMin(n *Node) *Node {
	// Traverse to the minimum child
	if n.leaf != nil {
		return n
	}
	if len(n.edges) > 0 {
		// Add all the other edges to the stack (the min node will be added as
		// we recurse)
		i.stack = append(i.stack, n.edges[1:])
		return i.recurseMin(n.edges[0].node)
	}
	// Shouldn't be possible
	return nil
}

// SeekLowerBound is used to seek the iterator to the smallest key that is
// greater or equal to the given key. There is no watch variant as it's hard to
// predict based on the radix structure which node(s) changes might affect the
// result.
func (i *Iterator) SeekLowerBound(key []byte) {
	// Wipe the stack. Unlike Prefix iteration, we need to build the stack as we
	// go because we need only a subset of edges of many nodes in the path to the
	// leaf with the lower bound.
	i.stack = []edges{}
	n := i.node
	search := key
	found := func(n *Node) {
		i.node = n
		i.stack = append(i.stack, edges{edge{node: n}})
	}
	for {
		// Compare current prefix with the search key's same-length prefix.
		var prefixCmp int
		if len(n.prefix) < len(search) {
			prefixCmp = bytes.Compare(n.prefix, search[0:len(n.prefix)])
		} else {
			prefixCmp = bytes.Compare(n.prefix, search)
		}
		if prefixCmp > 0 {
			// Prefix is larger, that means the lower bound is greater than the search
			// and from now on we need to follow the minimum path to the smallest
			// leaf under this subtree.
			n = i.recurseMin(n)
			if n != nil {
				found(n)
			}
			return
		}
		if prefixCmp < 0 {
			// Prefix is smaller than search prefix, that means there is no lower
			// bound
			i.node = nil
			return
		}
		// Prefix is equal, we are still heading for an exact match. If this is a
		// leaf we're done.
		if n.leaf != nil {
			if bytes.Compare(n.leaf.key, key) < 0 {
				i.node = nil
				return
			}
			found(n)
			return
		}
		// Consume the search prefix
		if len(n.prefix) > len(search) {
			search = []byte{}
		} else {
			search = search[len(n.prefix):]
		}
		// Otherwise, take the lower bound next edge.
		idx, lbNode := n.getLowerBoundEdge(search[0])
		if lbNode == nil {
			i.node = nil
			return
		}
		// Create stack edges for the all strictly higher edges in this node.
		if idx+1 < len(n.edges) {
			i.stack = append(i.stack, n.edges[idx+1:])
		}
		i.node = lbNode
		// Recurse
		n = lbNode
	}
}

// Next returns the next node in order
func (i *Iterator) Next() ([]byte, interface{}, bool) {
	// Initialize our stack if needed
	if i.stack == nil && i.node != nil {
		i.stack = []edges{
			{
				edge{node: i.node},
			},
		}
	}
	for len(i.stack) > 0 {
		// Inspect the last element of the stack
		n := len(i.stack)
		last := i.stack[n-1]
		elem := last[0].node
		// Update the stack
		if len(last) > 1 {
			i.stack[n-1] = last[1:]
		} else {
			i.stack = i.stack[:n-1]
		}
		// Push the edges onto the frontier
		if len(elem.edges) > 0 {
			i.stack = append(i.stack, elem.edges)
		}
		// Return the leaf values if any
		if elem.leaf != nil {
			return elem.leaf.key, elem.leaf.val, true
		}
	}
	return nil, nil, false
}
