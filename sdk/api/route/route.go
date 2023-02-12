package route

import (
	"net/http"
	"reflect"
	"strings"

	natsort "facette.io/natsort"
	prometheus "github.com/armon/go-metrics/prometheus"
)

// Route - gorilla mux route wrapper
// TODO: add validator/sanitizer functions.
// TODO: add setters
// TODO: add middleware functions
// TODO: type alias for []Route
type Route struct {
	name string
	// TODO: turn this into an array
	method      string
	path        string
	queries     map[string]string
	handlerFunc http.HandlerFunc
	middlewares []func(next http.HandlerFunc) http.HandlerFunc
	telemetry   *telemetry
}

// New returns an empty Route
// struct, ready for initialization
// through the Setter functions
func New() *Route {
	return &Route{
		queries:     make(map[string]string),
		handlerFunc: func(http.ResponseWriter, *http.Request) { panic("not implemented") },
		middlewares: make([]func(next http.HandlerFunc) http.HandlerFunc, 0),
		telemetry: &telemetry{
			gaugeDefinitions:   make([]prometheus.GaugeDefinition, 0),
			summaryDefinitions: make([]prometheus.SummaryDefinition, 0),
			counterDefinitions: make([]prometheus.CounterDefinition, 0),
		},
	}
}

// Equal answers whether v is equivalent to r.
// Always returns false if v is not a Route.
func (r Route) Equal(v interface{}) bool {
	r2, ok := v.(Route)
	if !ok {
		return false
	}
	if r.name != r2.name {
		return false
	}
	if r.method != r2.method {
		return false
	}
	if r.path != r2.path {
		return false
	}
	if !reflect.DeepEqual(r.queries, r2.queries) {
		return false
	}
	return true
}

//
// ──────────────────────────────────────────────────── I ──────────
//   :::::: S E T T E R : :  :   :    :     :        :          :
// ──────────────────────────────────────────────────────────────
//

// SetName sets the given value as name.
func (r *Route) SetName(s string) {
	r.name = s
}

// SetMethod sets the given value as method.
// TODO: add validator
// TODO: turn this into an array
func (r *Route) SetMethod(s string) {
	r.method = s
}

// SetPath sets the given value as path.
// TODO: add sanitizer
func (r *Route) SetPath(s string) {
	s = strings.Trim(s, "/")
	s = "/" + s
	r.path = s
}

// AppendQuery adds the given value to queries.
func (r *Route) AppendQuery(key, value string) {
	if r.queries == nil || len(r.queries) == 0 {
		r.queries = make(map[string]string)
	}
	_, ok := r.queries[key]
	if ok {
		return
	}
	r.queries[key] = value
}

// SetHandlerFunc sets the given value as handlerFunc.
func (r *Route) SetHandlerFunc(h http.HandlerFunc) {
	r.handlerFunc = h
}

// AppendMiddleware adds a middleware function
func (r *Route) AppendMiddleware(mw func(next http.HandlerFunc) http.HandlerFunc) {
	if r.middlewares == nil || len(r.middlewares) == 0 {
		r.middlewares = make([]func(next http.HandlerFunc) http.HandlerFunc, 0)
	}

	r.middlewares = append(r.middlewares, mw)
}

//
// ────────────────────────────────────────────────────── I ──────────
//   :::::: G E T T E R S : :  :   :    :     :        :          :
// ────────────────────────────────────────────────────────────────
//

// Name returns the value of name.
func (r Route) Name() string {
	return r.name
}

// Method returns the value of method.
func (r Route) Method() string {
	return r.method
}

// Path returns the value of path.
func (r Route) Path() string {
	return r.path
}

// Queries returns the value of queries.
func (r Route) Queries() []string {
	result := make([]string, 0)
	for key, value := range r.queries {
		result = append(result, key, value)
	}
	return result
}

// HandlerFunc returns the route handler after applying middlewares to it.
func (r Route) HandlerFunc() http.HandlerFunc {
	result := r.handlerFunc
	if middlewares := r.middlewares; middlewares != nil && len(middlewares) > 0 {
		for _, middleware := range middlewares {
			result = middleware(result)
		}
	}
	return result
}

//
// ──────────────────────────────────────────────── I ──────────
//   :::::: S O R T : :  :   :    :     :        :          :
// ──────────────────────────────────────────────────────────
//

// ByPath enables sorting of routes based on path
// usage:
type ByPath []Route

// Len implements sort interface
func (s ByPath) Len() int { return len(s) }

// Less implements sort interface
func (s ByPath) Less(a, b int) bool { return natsort.Compare(s[a].Path(), s[b].Path()) }

// Swap implements sort interface
func (s ByPath) Swap(a, b int) { s[a], s[b] = s[b], s[a] }

// ByName enables sorting routes based on their name
type ByName []Route

// Len implements sort interface
func (s ByName) Len() int { return len(s) }

// Less implements sort interface
func (s ByName) Less(a, b int) bool { return natsort.Compare(s[a].Name(), s[b].Name()) }

// Swap implements sort interface
func (s ByName) Swap(a, b int) { s[a], s[b] = s[b], s[a] }

// ByMethod enables sorting routes based on http verb (method)
type ByMethod []Route

// Len implements sort interface
func (s ByMethod) Len() int { return len(s) }

// Less implements sort interface
func (s ByMethod) Less(a, b int) bool { return natsort.Compare(s[a].Method(), s[b].Method()) }

// Swap implements sort interface
func (s ByMethod) Swap(a, b int) { s[a], s[b] = s[b], s[a] }
