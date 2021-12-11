package ews

import (
	"context"
	"fmt"

	"github.com/anfilat/go-ews/internal/base"
	"github.com/anfilat/go-ews/internal/errors"
	"github.com/anfilat/go-ews/internal/ews"
)

//nolint:unparam
func execute(_ context.Context, sd *ews.ServiceData, request base.Request) (interface{}, error) {
	if err := sd.Validate(); err != nil {
		return nil, err
	}
	if err := validateMinimumRequiredServerVersion(sd, request); err != nil {
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

func validateMinimumRequiredServerVersion(sd *ews.ServiceData, request base.Request) error {
	if sd.Version < request.GetMinimumRequiredServerVersion() {
		return errors.NewValidateError(fmt.Sprintf(
			"the service request %s is only valid for Exchange version %v or later.",
			request.GetXmlElementName(),
			request.GetMinimumRequiredServerVersion()))
	}
	return nil
}
