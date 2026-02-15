package api

import (
	"github.com/fedotovmax/green-api-test/internal/adapters/clients/greenapi"
	"github.com/go-chi/chi/v5"
)

type controller struct {
	greenApi *greenapi.Client
}

func New(greenApi *greenapi.Client) *controller {
	return &controller{
		greenApi: greenApi,
	}
}

func (c *controller) Register(router chi.Router) {
	//todo:
}
