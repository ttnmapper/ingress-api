package main

import (
	"github.com/766b/chi-prometheus"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Routes() *chi.Mux {
	router := chi.NewRouter()

	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.RealIP,
		middleware.Logger,
		middleware.Compress(5),
		middleware.StripSlashes,
		middleware.Recoverer,
		chiprometheus.NewMiddleware("ttnmapper-ingress-api", 0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9, 1, 1.5, 2, 5, 10, 100, 1000, 10000),
	)

	router.Handle("/metrics", promhttp.Handler())

	router.Route("/v1", func(r chi.Router) {
		r.Mount("/ttn", TtnRoutes())
		r.Mount("/android", AndroidRoutes())
		r.Mount("/ios", IosRoutes())
	})

	return router
}
