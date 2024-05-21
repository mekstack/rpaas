package serviceerror

import (
	"google.golang.org/grpc/status"
)

var (
	InvalidDomainError = status.Error(3, "invalid domain")
	InvalidHostError   = status.Error(3, "invalid host")
	InvalidPortError   = status.Error(3, "invalid port")
)
