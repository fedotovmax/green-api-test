package greenapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/fedotovmax/green-api-test/internal/domain"
	"github.com/fedotovmax/green-api-test/internal/keys"
)

func (gapi *Client) GetSettings(ctx context.Context, instanceID, tokenAPI string) (
	*domain.InstanceSettings, error) {

	const op = "adapters.clients.greenapi.GetSettings"

	url := buildURL(gapi.config.URL, instanceID, tokenAPI, GetSettingsMethod)

	req, err := http.NewRequestWithContext(ctx, keys.HTTPMethodGet, url, nil)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	req.Header.Set(keys.HeaderContentType, keys.ContentTypeJSON)

	resp, err := gapi.httpClient.Do(req)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {

		responseBodyBytes, err := io.ReadAll(resp.Body)

		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		var instanceSettings domain.InstanceSettings

		err = json.Unmarshal(responseBodyBytes, &instanceSettings)

		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		return &instanceSettings, nil

	}

	return nil, fmt.Errorf("%s: %w", op, ErrGreenAPIError)

}
