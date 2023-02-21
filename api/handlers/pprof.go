package handlers

import (
	"net/http"
	"net/http/pprof"

	middlewares "github.com/da-moon/podinfo/api/middlewares"
	registry "github.com/da-moon/podinfo/api/registry"
	route "github.com/da-moon/podinfo/sdk/api/route"
)

// debug enables profiling endpoints.
// TODO: enable with dev flag ?
// TODO: change namespace ?
func debug() {
	index := func() {
		name := "debug-index"
		path := "/pprof"
		r := route.New()
		r.SetName(name)
		r.SetPath(path)
		r.SetMethod(http.MethodGet)
		r.SetHandlerFunc(pprof.Index)
		log.Info("Adding log middleware for '%s' handler at '%s'", name, path)
		r.AppendMiddleware(middlewares.Log(log))
		registry.Register("/debug", *r)
	}
	allocs := func() {
		name := "debug-allocs"
		path := "/allocs"
		r := route.New()
		r.SetName(name)
		r.SetPath(path)
		r.SetMethod(http.MethodGet)
		r.SetHandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			pprof.Handler("allocs").ServeHTTP(w, r)
			return
		})
		log.Info("Adding log middleware for '%s' handler at '%s'", name, path)
		r.AppendMiddleware(middlewares.Log(log))
		registry.Register("/debug", *r)
	}
	block := func() {
		name := "debug-block"
		path := "/block"
		r := route.New()
		r.SetName(name)
		r.SetPath(path)
		r.SetMethod(http.MethodGet)
		r.SetHandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			pprof.Handler("block").ServeHTTP(w, r)
			return
		})
		log.Info("Adding log middleware for '%s' handler at '%s'", name, path)
		r.AppendMiddleware(middlewares.Log(log))
		registry.Register("/debug", *r)
	}
	cmdline := func() {
		name := "debug-cmdline"
		path := "/cmdline"
		r := route.New()
		r.SetName(name)
		r.SetPath(path)
		r.SetMethod(http.MethodGet)
		r.SetHandlerFunc(pprof.Cmdline)
		log.Info("Adding log middleware for '%s' handler at '%s'", name, path)
		r.AppendMiddleware(middlewares.Log(log))
		registry.Register("/debug", *r)
	}
	goroutine := func() {
		name := "debug-goroutine"
		path := "/goroutine"
		r := route.New()
		r.SetName(name)
		r.SetPath(path)
		r.SetMethod(http.MethodGet)
		r.SetHandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			pprof.Handler("goroutine").ServeHTTP(w, r)
			return
		})
		log.Info("Adding log middleware for '%s' handler at '%s'", name, path)
		r.AppendMiddleware(middlewares.Log(log))
		registry.Register("/debug", *r)
	}
	heap := func() {
		name := "debug-heap"
		path := "/heap"
		r := route.New()
		r.SetName(name)
		r.SetPath(path)
		r.SetMethod(http.MethodGet)
		r.SetHandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			pprof.Handler("heap").ServeHTTP(w, r)
			return
		})
		log.Info("Adding log middleware for '%s' handler at '%s'", name, path)
		r.AppendMiddleware(middlewares.Log(log))
		registry.Register("/debug", *r)
	}
	mutex := func() {
		name := "debug-mutex"
		path := "/mutex"
		r := route.New()
		r.SetName(name)
		r.SetPath(path)
		r.SetMethod(http.MethodGet)
		r.SetHandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			pprof.Handler("mutex").ServeHTTP(w, r)
			return
		})
		log.Info("Adding log middleware for '%s' handler at '%s'", name, path)
		r.AppendMiddleware(middlewares.Log(log))
		registry.Register("/debug", *r)
	}

	profile := func() {
		name := "debug-profile"
		path := "/profile"
		r := route.New()
		r.SetName(name)
		r.SetPath(path)
		r.SetMethod(http.MethodGet)
		r.SetHandlerFunc(pprof.Profile)
		log.Info("Adding log middleware for '%s' handler at '%s'", name, path)
		r.AppendMiddleware(middlewares.Log(log))
		registry.Register("/debug", *r)
	}
	threadcreate := func() {
		name := "debug-threadcreate"
		path := "/threadcreate"
		r := route.New()
		r.SetName(name)
		r.SetPath(path)
		r.SetMethod(http.MethodGet)
		r.SetHandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			pprof.Handler("threadcreate").ServeHTTP(w, r)
			return
		})
		log.Info("Adding log middleware for '%s' handler at '%s'", name, path)
		r.AppendMiddleware(middlewares.Log(log))
		registry.Register("/debug", *r)
	}
	symbol := func() {
		name := "debug-symbol"
		path := "/symbol"
		r := route.New()
		r.SetName(name)
		r.SetPath(path)
		r.SetMethod(http.MethodGet)
		r.SetHandlerFunc(pprof.Symbol)
		log.Info("Adding log middleware for '%s' handler at '%s'", name, path)
		r.AppendMiddleware(middlewares.Log(log))
		registry.Register("/debug", *r)
	}
	trace := func() {
		name := "debug-trace"
		path := "/trace"
		r := route.New()
		r.SetName(name)
		r.SetPath(path)
		r.SetMethod(http.MethodGet)
		r.SetHandlerFunc(pprof.Trace)
		log.Info("Adding log middleware for '%s' handler at '%s'", name, path)
		r.AppendMiddleware(middlewares.Log(log))
		registry.Register("/debug", *r)
	}
	// ────────────────────────────────────────────────────────────────────────────────
	index()
	allocs()
	block()
	cmdline()
	goroutine()
	heap()
	mutex()
	profile()
	threadcreate()
	symbol()
	trace()
}
