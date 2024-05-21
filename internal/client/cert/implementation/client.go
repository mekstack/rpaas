package certclientimpl

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log/slog"
	certclient "xds_server/internal/client/cert"

	pb "xds_server/pkg/cert"
)

type client struct {
	api    pb.CertClient
	conn   *grpc.ClientConn
	logger *slog.Logger
}

func New(certServerPort int, logger *slog.Logger) (certclient.Client, error) {
	c := &client{}
	c.logger = logger

	opts := make([]grpc.DialOption, 0, 1)
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.NewClient(fmt.Sprintf(":%d", certServerPort), opts...)
	if err != nil {
		return nil, err
	}

	c.conn = conn
	c.api = pb.NewCertClient(conn)

	return c, nil
}
