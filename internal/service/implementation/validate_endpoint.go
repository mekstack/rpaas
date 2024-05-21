package serviceimpl

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc/status"
	xdsmodels "xds_server/internal/models"
	serviceerror "xds_server/internal/service/errors"
)

const (
	domainField = "Domain"
	hostField   = "Host"
	portField   = "Port"
)

func (s *service) validateEndpoint(domainEp *xdsmodels.DomainEndpoint) error {
	if err := s.v.Struct(domainEp); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, err := range validationErrors {
			switch err.StructField() {
			case domainField:
				fmt.Println(serviceerror.InvalidDomainError)
				return serviceerror.InvalidDomainError
			case hostField:
				fmt.Println(serviceerror.InvalidHostError)
				return serviceerror.InvalidHostError
			case portField:
				fmt.Println(serviceerror.InvalidPortError)
				return serviceerror.InvalidPortError
			default:
				return status.Error(3, fmt.Sprintf("invalid data: %s", err.Error()))
			}
		}
	}

	return nil
}
