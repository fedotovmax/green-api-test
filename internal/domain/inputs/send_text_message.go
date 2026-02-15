package inputs

import (
	"errors"
	"fmt"

	"github.com/fedotovmax/green-api-test/internal/validation"
)

type SendTextMessage struct {
	ChatID  string `json:"chatId"`
	Message string `json:"message"`
}

func (i *SendTextMessage) Validate() error {
	var validationErrors []error

	err := validation.EmptyString(i.ChatID)

	if err != nil {
		validationErrors = append(validationErrors, fmt.Errorf("%s: %w", "ChatID", err))
	}

	err = validation.EmptyString(i.Message)

	if err != nil {
		validationErrors = append(validationErrors, fmt.Errorf("%s: %w", "Message", err))
	}

	return errors.Join(validationErrors...)
}
