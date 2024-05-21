package allroutesclientimpl

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log/slog"
	allroutesclient "xds_server/internal/client/routes"
	pb "xds_server/pkg/route_server"
)

type client struct {
	conn   *grpc.ClientConn
	api    pb.RouteClient
	logger *slog.Logger
}

func New(port int, logger *slog.Logger) (allroutesclient.Client, error) {
	c := &client{}
	c.logger = logger
	conn, err := grpc.NewClient(fmt.Sprintf(":%d", port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	c.conn = conn

	c.api = pb.NewRouteClient(conn)

	return c, nil
}
