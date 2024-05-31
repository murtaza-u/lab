package client

import (
	"context"
	pb "dummy/proto/gen/go"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

func Init() error {
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	client := pb.NewDummyClient(conn)
	ctx := context.Background()

	_, err = client.UnaryPing(ctx, &pb.Void{})
	if err != nil {
		return fmt.Errorf("UnaryPing: %s", err.Error())
	}
	fmt.Println("---------------------------------------------------")

	ss, err := client.ServerStreamPing(ctx, &pb.Void{})
	if err != nil {
		return fmt.Errorf("ServerStreamPing: %s", err.Error())
	}
	for {
		_, err := ss.Recv()
		if err != nil {
			log.Println("ServerStreamPing: connection ended by server")
			break
		}
		log.Println("ServerStreamPing: received message")
	}
	fmt.Println("---------------------------------------------------")

	cs, err := client.ClientStreamPing(ctx)
	if err != nil {
		return fmt.Errorf("ClientStreamPing: %s", err.Error())
	}
	for i := 0; i < 3; i++ {
		err := cs.Send(&pb.Void{})
		if err != nil {
			return fmt.Errorf("ClientStreamPing: %s", err.Error())
		}
		log.Println("ClientStreamPing: sent message")
	}
	return cs.CloseSend()
}
