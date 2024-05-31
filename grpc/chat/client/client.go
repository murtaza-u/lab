package client

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	pb "github.com/murtaza-u/chat/gen/go"
	"google.golang.org/grpc"
)

func Start() error {
	conn, err := grpc.Dial(":3000", grpc.WithInsecure())
	if err != nil {
		return err
	}

	client := pb.NewChatServiceClient(conn)

	ctx := context.Background()
	stream, err := client.Send(ctx)
	if err != nil {
		return err
	}

	go func(stream pb.ChatService_SendClient) {
		for {
			msg, err := stream.Recv()
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("\n> %s\n", msg.GetText())
		}
	}(stream)

	for {
		msg, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		msg = strings.TrimSpace(msg)
		if msg == "" {
			continue
		}

		err := stream.Send(&pb.Msg{Text: msg})
		if err != nil {
			return err
		}
	}
}
