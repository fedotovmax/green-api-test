package greenapi

import (
	"log/slog"
	"net/http"
)

type greenAPI struct {
	httpClient *http.Client
	log        *slog.Logger
	apiURL     string
}

func New(log *slog.Logger, greenApiURL string) *greenAPI {
	return &greenAPI{
		httpClient: &http.Client{},
		apiURL:     greenApiURL,
		log:        log,
	}
}
