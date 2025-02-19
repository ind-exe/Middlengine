// Package middlengine provides a simple middleware engine for chaining
// HTTP middleware functions in a Go web server.
package middlengine

import (
	"net/http"
)

// Middleware defines a function that takes an http.Handler and returns a new http.Handler.
// This allows you to add additional functionality (such as logging, authentication, etc.)
// to your HTTP handler.
type Middleware func(http.Handler) http.Handler

// Engine manages a stack of middleware functions and a base HTTP handler.
// It constructs a middleware chain that processes requests in the specified order.
type Engine struct {
	// Handler is the base HTTP handler that will be wrapped by the middleware chain.
	// This must be set before calling CreateChain.
	Handler http.Handler
	// Middlewares holds a slice of Middleware functions to be applied.
	// The order of insertion determines the order of middleware execution.
	Middlewares []Middleware
}

// NewEngine returns a new instance of Engine.
func NewEngine(handler http.Handler) *Engine {
	return &Engine{Handler: handler}
}

// Use appends a Middleware to the Engine's middleware stack.
// Middlewares are applied in the order they are Used (the first Used is the outermost).
func (e *Engine) Use(middleware Middleware) {
	e.Middlewares = append(e.Middlewares, middleware)
}

// CreateChain builds the middleware chain by wrapping the base handler.
// It iterates through the Middlewares slice in reverse order, ensuring that the first
// middleware added becomes the outermost layer.
func (e *Engine) CreateChain() {
	// Ensure that a base handler has been set to avoid runtime errors.
	if e.Handler == nil {
		panic("base handler is nil: set a valid http.Handler before creating the middleware chain")
	}
	// Wrap the handler with each middleware in reverse order.
	for i := len(e.Middlewares) - 1; i >= 0; i-- {
		e.Handler = e.Middlewares[i](e.Handler)
	}
}

// ServeHTTP satisfies the http.Handler interface.
// It delegates the HTTP request to the fully composed middleware chain.
func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	e.Handler.ServeHTTP(w, r)
}
