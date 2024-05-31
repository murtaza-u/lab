package server

import (
	"context"
	pb "dummy/proto/gen/go"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedDummyServer
}

func (s *server) UnaryPing(ctx context.Context, _ *pb.Void) (*pb.Void, error) {
	log.Println("Hit UnaryPing")
	return &pb.Void{}, nil
}

func (s *server) ServerStreamPing(in *pb.Void, stream pb.Dummy_ServerStreamPingServer) error {
	log.Println("Hit StreamPingServer")
	for i := 0; i < 3; i++ {
		err := stream.Send(&pb.Void{})
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *server) ClientStreamPing(stream pb.Dummy_ClientStreamPingServer) error {
	log.Println("Hit ClientStreamPing")
	for {
		_, err := stream.Recv()
		if err != nil {
			log.Println("ClientStreamPing: connection closed by client")
			return nil
		}
		log.Println("ClientStreamPing: received a message")
	}
}

func unaryInterceptor(
	ctx context.Context,
	req any,
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (any, error) {

	log.Printf(
		"Hit unary interceptor: redirecting to %s\n",
		info.FullMethod,
	)
	return handler(ctx, req)
}

func streamInterceptor(
	srv any,
	ss grpc.ServerStream,
	info *grpc.StreamServerInfo,
	handler grpc.StreamHandler) error {

	log.Printf(
		"Hit stream interceptor: redirecting to %s\n",
		info.FullMethod,
	)
	return handler(srv, ss)
}

func Init() error {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		return err
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(unaryInterceptor),
		grpc.StreamInterceptor(streamInterceptor),
	)
	pb.RegisterDummyServer(s, &server{})
	return s.Serve(ln)
}
