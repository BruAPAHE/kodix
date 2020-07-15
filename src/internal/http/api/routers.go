package api

import (
	"github.com/gorilla/mux"
	v1 "kodix/src/internal/http/api/handlers/v1"
	"kodix/src/internal/http/api/middlewares"
	"net/http"
)

type Router struct {
	Handler *mux.Router
}

func NewRouter() *Router {
	r := mux.NewRouter()
	r.StrictSlash(true)

	return &Router{Handler: r}
}

// Router List ...
func (r *Router) RegisterRoutes() {
	// Auto ...
	r.Handler.HandleFunc("/v1/auto/{id:[0-9]+}",
		middlewares.ApplyMiddleware(
			v1.GetAutoById(),
			middlewares.HeadersMiddleware,
		),
	).Methods(http.MethodGet)

	r.Handler.HandleFunc("/v1/auto/",
		middlewares.ApplyMiddleware(
			v1.GetAllAuto(),
			middlewares.HeadersMiddleware,
		),
	).Methods(http.MethodGet)

	r.Handler.HandleFunc("/v1/auto/",
		middlewares.ApplyMiddleware(
			v1.CreateUser(),
			middlewares.HeadersMiddleware,
		),
	).Methods(http.MethodPost)

	r.Handler.HandleFunc("/v1/auto/{id:[0-9]+}",
		middlewares.ApplyMiddleware(
			v1.UpdateAutoById(),
			middlewares.HeadersMiddleware,
		),
	).Methods(http.MethodPatch)

	r.Handler.HandleFunc("/v1/auto/{id:[0-9]+}",
		middlewares.ApplyMiddleware(
			v1.DeleteAutoById(),
			middlewares.HeadersMiddleware,
		),
	).Methods(http.MethodDelete)
}
