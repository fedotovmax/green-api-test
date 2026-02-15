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

func (gapi *greenAPI) GetStateInstance(ctx context.Context, instanceID, tokenAPI string) (
	*domain.InstanceStateResponse, error,
) {
	const op = "adapters.clients.greenapi.GetStateInstance"

	url := buildURL(gapi.apiURL, instanceID, tokenAPI, GetStateInstanceMethod)

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

		var instanceState domain.InstanceStateResponse

		err = json.Unmarshal(responseBodyBytes, &instanceState)

		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		return &instanceState, nil

	}

	return nil, fmt.Errorf("%s: %w", op, ErrGreenAPIError)

}
