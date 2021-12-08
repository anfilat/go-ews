package ews

import (
	"context"

	"github.com/anfilat/go-ews/internal/requests"
)

//nolint:unparam
func execute(_ context.Context, service *ExchangeService, request requests.Request) (interface{}, error) {
	if err := service.validate(); err != nil {
		return nil, err
	}
	if err := request.Validate(); err != nil {
		return nil, err
	}

	service.ensureClient()

	return nil, nil
}
