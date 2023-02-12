package pathmanager

import (
	"strings"
	"sync"

	iradix "github.com/da-moon/northern-labs-interview/internal/radix-tree/immutable"
)

// PathManager is a prefix searchable index of paths
type PathManager struct {
	l     sync.RWMutex
	paths *iradix.Tree
}

// New creates a new path manager
func New() *PathManager {
	return &PathManager{
		paths: iradix.New(),
	}
}

// AddPaths adds path to the paths list
func (p *PathManager) AddPaths(paths []string) {
	p.l.Lock()
	defer p.l.Unlock()
	txn := p.paths.Txn()
	for _, prefix := range paths {
		if prefix == "" {
			continue
		}
		// Exception ...
		var Exception bool
		if strings.HasPrefix(prefix, "!") {
			prefix = strings.TrimPrefix(prefix, "!")
			Exception = true
		}
		// We trim any trailing *, but we don't touch whether it is a trailing
		// slash or not since we want to be able to ignore prefixes that fully
		// specify a file
		txn.Insert([]byte(strings.TrimSuffix(prefix, "*")), Exception)
	}
	p.paths = txn.Commit()
}

// RemovePaths removes paths from the paths list
func (p *PathManager) RemovePaths(paths []string) {
	p.l.Lock()
	defer p.l.Unlock()
	txn := p.paths.Txn()
	for _, prefix := range paths {
		if prefix == "" {
			continue
		}
		// Exceptions aren't stored with the leading ! so strip it
		if strings.HasPrefix(prefix, "!") {
			prefix = strings.TrimPrefix(prefix, "!")
		}
		// We trim any trailing *, but we don't touch whether it is a trailing
		// slash or not since we want to be able to ignore prefixes that fully
		// specify a file
		txn.Delete([]byte(strings.TrimSuffix(prefix, "*")))
	}
	p.paths = txn.Commit()
}

// RemovePathPrefix removes all paths with the given prefix
func (p *PathManager) RemovePathPrefix(prefix string) {
	p.l.Lock()
	defer p.l.Unlock()
	// We trim any trailing *, but we don't touch whether it is a trailing
	// slash or not since we want to be able to ignore prefixes that fully
	// specify a file
	p.paths, _ = p.paths.DeletePrefix([]byte(strings.TrimSuffix(prefix, "*")))
}

// Len returns the number of paths
func (p *PathManager) Len() int {
	return p.paths.Len()
}

// Paths returns the path list
func (p *PathManager) Paths() []string {
	p.l.RLock()
	defer p.l.RUnlock()
	paths := make([]string, 0, p.paths.Len())
	walkFn := func(k []byte, v interface{}) bool {
		paths = append(paths, string(k))
		return false
	}
	p.paths.Root().Walk(walkFn)
	return paths
}

// HasPath returns if the prefix for the path exists regardless if it is a path
// (ending with /) or a prefix for a leaf node
func (p *PathManager) HasPath(path string) bool {
	p.l.RLock()
	defer p.l.RUnlock()
	if _, ExceptionRaw, ok := p.paths.Root().LongestPrefix([]byte(path)); ok {
		// Exception ...
		var Exception bool
		if ExceptionRaw != nil {
			Exception = ExceptionRaw.(bool)
		}
		return !Exception
	}
	return false
}

// HasExactPath function returns if the longest match is an exact match for the
// full path
func (p *PathManager) HasExactPath(path string) bool {
	p.l.RLock()
	defer p.l.RUnlock()
	if val, ExceptionRaw, ok := p.paths.Root().LongestPrefix([]byte(path)); ok {
		// Exception ...
		var Exception bool
		if ExceptionRaw != nil {
			Exception = ExceptionRaw.(bool)
		}
		strVal := string(val)
		if strings.HasSuffix(strVal, "/") || strVal == path {
			return !Exception
		}
	}
	return false
}
