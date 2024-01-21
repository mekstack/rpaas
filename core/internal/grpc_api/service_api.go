package grpc_api

import (
	"github.com/mekstack/nataas/core/internal/controller"
)

type serviceApi struct {
	cnt *controller.Controller
}

func New(cnt controller.Controller) *serviceApi {
	return &serviceApi{
		cnt: &cnt,
	}
}
