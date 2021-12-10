package ews

import (
	"context"
	"fmt"

	"github.com/anfilat/go-ews/internal/ews"
	"github.com/anfilat/go-ews/internal/requests"
)

//nolint:unparam
func execute(_ context.Context, sd *ews.ServiceData, request requests.Request) (interface{}, error) {
	if err := sd.Validate(); err != nil {
		return nil, err
	}
	if err := request.Validate(); err != nil {
		return nil, err
	}

	sd.EnsureClient()

	writer := ews.NewRequestWriter(sd, request)
	buf, err := writer.WriteXML()
	fmt.Println(string(buf))
	fmt.Println(err)

	return nil, nil
}
