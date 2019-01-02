package main

import (
	"github.com/766b/chi-prometheus"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/prometheus/client_golang/prometheus"
)

func Routes() *chi.Mux {
	router := chi.NewRouter()

	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.DefaultCompress,
		middleware.RedirectSlashes,
		middleware.Recoverer,
		chiprometheus.NewMiddleware("ttnmapper-ingress-api"),
	)

	router.Handle("/metrics", prometheus.Handler())

	router.Route("/v1", func(r chi.Router) {
		r.Mount("/ttn", TtnRoutes())
		r.Mount("/android", AndroidRoutes())
	})

	return router
}
