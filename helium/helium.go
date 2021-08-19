package helium

import (
	"github.com/go-chi/chi"
	"ttnmapper-ingress-api/types"
)

type Context struct {
	PublishChannel chan types.TtnMapperUplinkMessage
}

func HeliumRoutes(publishChannel chan types.TtnMapperUplinkMessage) *chi.Mux {
	context := &Context{PublishChannel: publishChannel}

	router := chi.NewRouter()

	router.Post("/v1", context.PostHelium)
	router.Get("/v1", context.GetHelium)

	return router
}
