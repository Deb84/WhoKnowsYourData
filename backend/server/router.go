package server

import (
	"net/http"
	"whoknowsyourdata/domain"

	"github.com/go-chi/chi/v5"
)

// Router is used over chi to use ErrorHandler
type Router struct {
	*chi.Mux
	Log          domain.Logger
	errorHandler ErrorHandler
}

type RouteParams struct {
	Route string
	Param map[string]string
}

type AppHandlerFunc func(http.ResponseWriter, *http.Request) error

type HandlerFuncParam func(http.ResponseWriter, *http.Request, string) error

func NewRouter(log domain.Logger, errorHandler *ErrorHandler) *Router {
	return &Router{
		Mux:          chi.NewRouter(),
		Log:          log,
		errorHandler: *errorHandler,
	}
}

// HTTP methods with HandlerFunc wrapped by ErrorHandler

func (r *Router) Connect(route string, fn AppHandlerFunc) {
	r.Mux.Connect(route, r.errorHandler.Wrap(fn))
}

func (r *Router) Delete(route string, fn AppHandlerFunc) {
	r.Mux.Delete(route, r.errorHandler.Wrap(fn))
}

func (r *Router) Get(route string, fn AppHandlerFunc) {
	r.Mux.Get(route, r.errorHandler.Wrap(fn))
}

func (r *Router) Head(route string, fn AppHandlerFunc) {
	r.Mux.Head(route, r.errorHandler.Wrap(fn))
}

func (r *Router) Options(route string, fn AppHandlerFunc) {
	r.Mux.Options(route, r.errorHandler.Wrap(fn))
}

func (r *Router) Patch(route string, fn AppHandlerFunc) {
	r.Mux.Patch(route, r.errorHandler.Wrap(fn))
}

func (r *Router) Post(route string, fn AppHandlerFunc) {
	r.Mux.Post(route, r.errorHandler.Wrap(fn))
}

func (r *Router) Put(route string, fn AppHandlerFunc) {
	r.Mux.Put(route, r.errorHandler.Wrap(fn))
}

func (r *Router) Query(route string, fn AppHandlerFunc) {
	r.Mux.Query(route, r.errorHandler.Wrap(fn))
}

func (r *Router) Trace(route string, fn AppHandlerFunc) {
	r.Mux.Trace(route, r.errorHandler.Wrap(fn))
}
