package server

import (
	"log"
	"net"

	pb "github.com/murtaza-u/chat/gen/go"
	"google.golang.org/grpc"
)

const port = "3000"

type service struct {
	clients []pb.ChatService_SendServer

	pb.UnimplementedChatServiceServer
}

func (s *service) broadcast(msg string) {
	for _, c := range s.clients {
		log.Printf("sending message %s to a client\n", msg)
		c.Send(&pb.Msg{Text: msg})
	}
}

func (s *service) remove(client pb.ChatService_SendServer) {
	for i := 0; i < len(s.clients); i++ {
		if client == s.clients[i] {
			s.clients = append(s.clients[:i], s.clients[i+1:]...)
			return
		}
	}
}

func (s *service) Send(stream pb.ChatService_SendServer) error {
	s.clients = append(s.clients, stream)

	for {
		msg, err := stream.Recv()
		if err != nil {
			s.remove(stream)
			return nil
		}

		s.broadcast(msg.GetText())
	}
}

func Start() error {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	s := grpc.NewServer()
	pb.RegisterChatServiceServer(s, &service{})
	return s.Serve(lis)
}
