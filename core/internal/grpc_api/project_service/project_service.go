package project_service

import (
	"context"
	"github.com/mekstack/nataas/core/internal/controller"
	proto "github.com/mekstack/nataas/core/proto/pb"
	"google.golang.org/grpc"
)

type ProjectApi struct {
	controller *controller.Controller
	proto.UnimplementedProjectServiceServer
}

func (p *ProjectApi) GetProjectInfo(ctx context.Context, request *proto.GetProjectInfoRequest) (*proto.GetProjectInfoResponse, error) {
	response, err := (*p.controller).GetProjectInfo(ctx, request.Code)
	if err != nil {
		return nil, err
	}
	return &proto.GetProjectInfoResponse{Project: response}, nil
}
func (p *ProjectApi) AddProjectInfo(ctx context.Context, request *proto.AddProjectInfoRequest) (*proto.AddProjectInfoResponse, error) {
	response, err := (*p.controller).AddProjectInfo(ctx, request.Code, request.Route)
	if err != nil {
		return nil, err
	}
	return &proto.AddProjectInfoResponse{Project: response}, nil
}

func Register(server *grpc.Server, controller controller.Controller) {
	proto.RegisterProjectServiceServer(server, &ProjectApi{
		controller: &controller,
	})
}
