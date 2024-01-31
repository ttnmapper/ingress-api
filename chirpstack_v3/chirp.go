package chirpstack_v3

import (
	"github.com/go-chi/chi"
	_ "time"
	"ttnmapper-ingress-api/types"
)

type Context struct {
	PublishChannel chan types.TtnMapperUplinkMessage
}

func ChirpRoutes(publishChannel chan types.TtnMapperUplinkMessage) *chi.Mux {
	context := &Context{PublishChannel: publishChannel}

	router := chi.NewRouter()

	router.Post("/events", context.PostChirpV3Event)
	return router
}
