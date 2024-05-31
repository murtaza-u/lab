package server

import (
	"context"
	"log"
	"net"
	"path/filepath"

	pb "github.com/murtaza-u/arith/gen/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
)

type server struct {
	pb.UnimplementedArithServiceServer
}

func Run() error {
	l, err := net.Listen("tcp", ":10392")
	if err != nil {
		return err
	}

	defer l.Close()

	creds, err := credentials.NewServerTLSFromFile(
		filepath.Join("certs", "server.pem"),
		filepath.Join("certs", "server.key"),
	)
	if err != nil {
		return err
	}

	s := grpc.NewServer(grpc.Creds(creds))
	pb.RegisterArithServiceServer(s, &server{})

	return s.Serve(l)
}

func (s *server) Compute(c context.Context, r *pb.Request) (*pb.Response, error) {
	// emulating heavy workload
	// time.Sleep(time.Second * 5)

	x, y, op := r.GetX(), r.GetY(), r.GetOp()
	var err error
	var res int64

	switch op {
	case pb.Operator_OPERATOR_UNDEFINED:
		err = status.Error(codes.InvalidArgument, "undefined operator")

	case pb.Operator_OPERATOR_PLUS:
		res = x + y

	case pb.Operator_OPERATOR_MINUS:
		res = x - y

	case pb.Operator_OPERATOR_MULTIPLY:
		res = x * y

	case pb.Operator_OPERATOR_DIVIDE:
		res = x / y

	default:
		err = status.Errorf(codes.Unimplemented, "unimplemented operator: %v", op)
	}

	if err != nil {
		log.Printf("An error occured: %s\n", err.Error())
		return nil, err
	}

	if c.Err() == context.DeadlineExceeded {
		log.Println("Stopping RPC: context dealine exceeded")
		return nil, status.Error(codes.DeadlineExceeded, c.Err().Error())
	}

	if c.Err() == context.Canceled {
		log.Println("Stopping RPC: request cancelled")
		return nil, status.Error(codes.Canceled, c.Err().Error())
	}

	log.Printf("computed value: %d\n", res)
	return &pb.Response{Value: res}, nil
}
