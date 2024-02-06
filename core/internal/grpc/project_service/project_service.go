package project_service

import (
	"context"

	proto "github.com/mekstack/nataas/core/proto/pb"

	"google.golang.org/grpc"
)

type ProjectController interface {
	GetProject(context.Context, uint32) (*proto.Project, error)
	AddRouteToProject(context.Context, uint32, *proto.Route) (*proto.Project, error)
}

type service struct {
	proto.UnimplementedProjectServiceServer
	projectController ProjectController
}

func Register(server *grpc.Server, controller ProjectController) {
	proto.RegisterProjectServiceServer(server, &service{
		projectController: controller,
	})
}

func (p *service) GetProject(ctx context.Context, request *proto.GetProjectRequest) (*proto.GetProjectResponse, error) {
	project, err := p.projectController.GetProject(ctx, request.Code)
	if err != nil {
		return nil, err
	}
	return &proto.GetProjectResponse{Project: project}, nil
}

func (p *service) AddRouteToProject(ctx context.Context, request *proto.AddRouteToProjectRequest) (*proto.AddRouteToProjectResponse, error) {
	project, err := p.projectController.AddRouteToProject(ctx, request.Code, request.Route)
	if err != nil {
		return nil, err
	}
	return &proto.AddRouteToProjectResponse{Project: project}, nil
}
