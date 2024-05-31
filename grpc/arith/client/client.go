package client

import (
	"context"
	"fmt"
	"path/filepath"

	pb "github.com/murtaza-u/arith/gen/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
)

func Add(x, y int64) error {
	creds, err := credentials.NewClientTLSFromFile(
		filepath.Join("certs", "server.pem"), "server",
	)
	if err != nil {
		return err
	}

	conn, err := grpc.Dial(
		"3.109.50.164:1449",
		grpc.WithTransportCredentials(creds),
	)
	if err != nil {
		return err
	}

	defer conn.Close()

	client := pb.NewArithServiceClient(conn)

	ctx := context.Background()

	// ctx, cancel := context.WithTimeout(context.Background(), time.Second*4)
	// defer cancel()

	resp, err := client.Compute(ctx, &pb.Request{
		X: x, Y: y, Op: pb.Operator_OPERATOR_PLUS,
	})

	if err != nil {
		if errStatus, ok := status.FromError(err); ok {
			return status.Error(errStatus.Code(), errStatus.Message())
		}

		return err
	}

	val := resp.GetValue()
	fmt.Printf("value: %d\n", val)
	return nil
}
