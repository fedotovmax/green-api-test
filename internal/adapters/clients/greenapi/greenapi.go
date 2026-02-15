package greenapi

import (
	"log/slog"
	"net/http"

	"github.com/fedotovmax/green-api-test/internal/config"
)

type Client struct {
	httpClient *http.Client
	log        *slog.Logger
	config     *config.GreenAPIConfig
}

func New(log *slog.Logger, config *config.GreenAPIConfig) *Client {
	return &Client{
		httpClient: &http.Client{},
		config:     config,
		log:        log,
	}
}
