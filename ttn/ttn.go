package ttn

import (
	"github.com/go-chi/chi"
	"ttnmapper-ingress-api/types"
)

type Context struct {
	PublishChannel chan types.TtnMapperUplinkMessage
}

func TtnRoutes(publishChannel chan types.TtnMapperUplinkMessage) *chi.Mux {
	context := &Context{PublishChannel: publishChannel}

	router := chi.NewRouter()

	router.Post("/v2", context.PostTtnV2)
	router.Get("/v2", context.GetTtnV2)

	return router
}
