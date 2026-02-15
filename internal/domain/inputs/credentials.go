package inputs

import (
	"errors"
	"fmt"

	"github.com/fedotovmax/green-api-test/internal/validation"
)

type Credentials struct {
	APIToken   string
	InstanceID string
}

func (i *Credentials) Validate() error {

	var validationErrors []error

	err := validation.EmptyString(i.InstanceID)

	if err != nil {
		validationErrors = append(validationErrors, fmt.Errorf("%s: %w", "InstanceID", err))
	}

	err = validation.EmptyString(i.APIToken)

	if err != nil {
		validationErrors = append(validationErrors, fmt.Errorf("%s: %w", "APIToken", err))
	}

	return errors.Join(validationErrors...)

}
