package inputs

import (
	"errors"
	"fmt"

	"github.com/fedotovmax/green-api-test/internal/validation"
)

type SendFile struct {
	ChatID   string `json:"chatId"`
	FileURL  string `json:"urlFile"`
	FileName string `json:"fileName"`
}

func (i *SendFile) Validate() error {
	var validationErrors []error

	err := validation.EmptyString(i.ChatID)

	if err != nil {
		validationErrors = append(validationErrors, fmt.Errorf("%s: %w", "ChatID", err))
	}

	err = validation.EmptyString(i.FileName)

	if err != nil {
		validationErrors = append(validationErrors, fmt.Errorf("%s: %w", "FileName", err))
	}

	_, err = validation.IsURI(i.FileURL)

	if err != nil {
		validationErrors = append(validationErrors, fmt.Errorf("%s: %w", "FileURL", err))
	}

	return errors.Join(validationErrors...)

}
