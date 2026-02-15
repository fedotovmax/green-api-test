package greenapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/fedotovmax/green-api-test/internal/domain"
	"github.com/fedotovmax/green-api-test/internal/domain/inputs"
	"github.com/fedotovmax/green-api-test/internal/keys"
)

func (gapi *Client) SendFileByUrl(
	ctx context.Context,
	instanceID, tokenAPI string,
	in *inputs.SendFile,
) (
	*domain.NewMessage, error,
) {
	const op = "adapters.clients.greenapi.SendFileByUrl"

	url := buildURL(gapi.config.URL, instanceID, tokenAPI, SendFileByUrlMethod)

	inputBytes, err := json.Marshal(in)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	req, err := http.NewRequestWithContext(ctx, keys.HTTPMethodPost, url, bytes.NewReader(inputBytes))

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

		var newMessage domain.NewMessage

		err = json.Unmarshal(responseBodyBytes, &newMessage)

		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		return &newMessage, nil

	}

	return nil, fmt.Errorf("%s: %w", op, ErrGreenAPIError)
}
