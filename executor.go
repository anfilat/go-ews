package ews

import (
	"context"
	"fmt"

	"github.com/anfilat/go-ews/internal/ews"
	"github.com/anfilat/go-ews/internal/requests"
)

//nolint:unparam
func execute(_ context.Context, service *ews.Data, request requests.Request) (interface{}, error) {
	if err := service.Validate(); err != nil {
		return nil, err
	}
	if err := request.Validate(); err != nil {
		return nil, err
	}

	service.EnsureClient()

	writer := ews.NewRequestWriter(service, request)
	buf, err := writer.WriteXML()
	fmt.Println(string(buf))
	fmt.Println(err)

	return nil, nil
}
