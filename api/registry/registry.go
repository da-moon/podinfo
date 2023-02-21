package registry

import (
	"sync"

	route "github.com/da-moon/podinfo/sdk/api/route"
)

var (
	// routes collects all endpoints
	// nolint:gochecknoglobals // this is private and only methods in the same package can access it
	routes = make(map[string][]route.Route)
	// lock guards the routes array from
	// concurrent modifications.
	// nolint:gochecknoglobals // this is private and only methods in the same package can access it
	mutex sync.RWMutex
)

// Register adds a new route to the registry
func Register(prefix string, h route.Route) {
	mutex.Lock()
	defer mutex.Unlock()
	if prefix == "" {
		prefix = "/"
	}
	subroutes, ok := routes[prefix]
	if !ok || subroutes == nil {
		subroutes = make([]route.Route, 0)
	}
	subroutes = append(subroutes, h)
	routes[prefix] = subroutes
}

// Dispense returns all routers
func Dispense() (*route.Collection, error) {
	mutex.RLock()
	defer mutex.RUnlock()
	result := route.NewCollection()
	for k, v := range routes {
		for _, vv := range v {
			result.AppendRoute(k, vv)
		}
	}
	return result, nil
}

// Prefixes returns an array of subrouter
// prefixes , e.g "/v1"
func Prefixes() []string {
	mutex.RLock()
	defer mutex.RUnlock()
	result := make([]string, 0)
	for k := range routes {
		result = append(result, k)
	}
	return result
}
