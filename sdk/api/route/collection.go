package route

import (
	"fmt"
	"io"
	"sort"
	"strings"
	"sync"

	prometheus "github.com/armon/go-metrics/prometheus"
	mux "github.com/gorilla/mux"
)

// Collection helps with aggeragating routes.
// it also enables for a route struct to be self documenting
type Collection struct {
	mutex  sync.RWMutex
	routes map[string][]Route
}

// NewCollection creates struct
// that aggregates all routes in a
// single api version
// func NewCollection(routes []Route, opts ...CollectionOption) *Collection {
func NewCollection() *Collection {
	//nolint:gocritic // keeping this for potential future use
	// if routes == nil || len(routes) == 0 {
	// 	routes = make(map[string][]Route)
	// }
	//lint:gocritic
	result := &Collection{
		routes: make(map[string][]Route),
	}
	//nolint:gocritic // keeping this for potential future use
	// for _, opt := range opts {
	// 	opt(result)
	// }
	//lint:gocritic
	return result
}

// AppendRoute appends a route to the collection
func (c *Collection) AppendRoute(prefix string, r Route) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.routes == nil {
		c.routes = make(map[string][]Route)
	}
	if prefix == "" {
		prefix = "/"
	}
	routes, ok := c.routes[prefix]
	if !ok || routes == nil {
		routes = make([]Route, 0)
	}

	routes = append(routes, r)
	c.routes[prefix] = routes
}

// Synopsis returns a short description about all the routes in a
// single collection
func (c *Collection) Synopsis(w io.Writer, host string) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	if host != "" {
		fmt.Fprintf(w, "\n+ Listening On  : %s ", host)
	}
	for prefix, routes := range c.routes {
		if prefix == "/" {
			fmt.Fprintf(w, "\n+ Global API : ")
		} else {
			fmt.Fprintf(w, "\n+ API Version (Prefix) : %s ", prefix)
		}
		sort.Sort(ByPath(routes))
		fmt.Fprintf(w, "\n+ Routes:")
		for _, route := range routes {
			//nolint:gocritic // keeping this for potential future use
			// path := "/" + strings.Join([]string{strings.Trim(prefix, "/"), strings.Trim(route.Path(), "/")}, "/")
			//lint:gocritic
			path := strings.Trim(prefix, "/") + "/" + strings.Trim(route.Path(), "/")
			if host != "" {
				//nolint:gocritic // keeping this for potential future use
				// path = strings.Join([]string{strings.Trim(host, "/"), strings.Trim(path, "/")}, "/")
				//lint:gocritic
				path = strings.Trim(host, "/") + "/" + strings.Trim(path, "/")
			}
			fmt.Fprintf(w, "\n[ %s ] %s\t%s", route.Method(), route.Name(), path)
		}
		fmt.Fprintf(w, "\n\n")
	}
}

// Router returns a new Gorilla Mux Wrapper
func (c *Collection) Router() *mux.Router {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	router := mux.NewRouter().PathPrefix("/").Subrouter()
	if c.routes == nil || len(c.routes) == 0 {
		return router
	}
	for prefix, routes := range c.routes {
		for _, route := range routes {
			router.
				PathPrefix(prefix).
				Methods(route.Method()).
				Path(route.Path()).
				Name(route.Name()).
				Handler(route.HandlerFunc()).
				Queries(route.Queries()...)
		}
	}
	return router
}

// GaugeDefinitions returns a list of all the gauge definitions
func (c *Collection) GaugeDefinitions() []prometheus.GaugeDefinition {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	var defs []prometheus.GaugeDefinition
	for _, routes := range c.routes {
		for _, route := range routes {
			defs = append(defs, route.GaugeDefinitions()...)
		}
	}
	return defs
}

// SummaryDefinitions returns a list of all the summary definitions
func (c *Collection) SummaryDefinitions() []prometheus.SummaryDefinition {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	var defs []prometheus.SummaryDefinition
	for _, routes := range c.routes {
		for _, route := range routes {
			defs = append(defs, route.SummaryDefinitions()...)
		}
	}
	return defs
}

// CounterDefinitions returns a list of all the counter definitions
func (c *Collection) CounterDefinitions() []prometheus.CounterDefinition {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	var defs []prometheus.CounterDefinition
	for _, routes := range c.routes {
		for _, route := range routes {
			defs = append(defs, route.CounterDefinitions()...)
		}
	}
	return defs
}
