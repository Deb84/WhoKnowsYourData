// Package server provides the HTTP server, router, routes and error handler
package server

import (
	"net/http"
)

func New(addr string, router http.Handler) *http.Server {

	server := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	return server
}
