// package mutable implements a mutable radix tree(http://en.wikipedia.org/wiki/Radix_tree).
// The package only provides a single `Tree` implementation, optimized for sparse nodes.
// As a radix tree, it provides the following:
//
//	-O(k) operations. In many cases, this can be faster than a hash table since
//	  the hash function is an O(k) operation, and hash tables have very poor cache locality.
//	- Minimum / Maximum value lookups
//	- Ordered iteration
//
// A tree supports using a transaction to batch multiple updates (insert, delete)
// in a more efficient manner than performing each operation one at a time.
package mutable
