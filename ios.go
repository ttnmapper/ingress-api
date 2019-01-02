package main

import "github.com/go-chi/chi"

func IosRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Post("/v2", PostAndroidV2) // same format as Android v2

	return router
}
