package main

import (
	"fmt"
	"log"
	"os"

	pb "github.com/murtaza-u/search/gen/go"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

func main() {
	req := &pb.SearchRequest{
		Query:         "protocol buffers",
		PageNumber:    1,
		ResultPerPage: 10,
		EngineId: &pb.SearchRequest_Duck{
			Duck: "duck_981",
		},
	}

	fmt.Println(req.GetBing())
	fmt.Println(req.GetGoogle())
	fmt.Println(req.GetDuck())

	any, err := anypb.New(req)
	if err != nil {
		log.Fatal(err)
	}

	status := &pb.ErrorStatus{
		Message: "An error occured",
		Details: []*anypb.Any{any},
	}

	fmt.Println(status.GetMessage())
	fmt.Println(status.GetDetails())

	// write wire-format encoding to file
	b, err := proto.Marshal(req)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile("request.bin", b, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
