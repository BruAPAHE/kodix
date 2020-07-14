package middlewares

import "net/http"

func HeadersMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		// next
		handler(writer, request)
	}
}
