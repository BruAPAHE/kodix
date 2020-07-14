package middlewares

import "net/http"

type Middleware func(http.HandlerFunc) http.HandlerFunc

// iterator of middleware
func ApplyMiddleware(handle http.HandlerFunc, middleware ...Middleware) http.HandlerFunc {
	if len(middleware) < 1 {
		return handle
	}
	wrapped := handle

	for i := len(middleware) - 1; i >= 0; i-- {
		wrapped = middleware[i](wrapped)
	}

	return wrapped
}
