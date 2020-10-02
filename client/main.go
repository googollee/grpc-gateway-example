package main

import (
	"client/servicepb"
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
)

type preRPC map[string]string

func (p preRPC) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	fmt.Println("pre rpc, uri:", uri)
	return map[string]string(p), nil
}

func (p preRPC) RequireTransportSecurity() bool {
	return false
}

func main() {
	ctx := context.Background()

	server := "localhost:50551"
	conn, err := grpc.Dial(server,
		grpc.WithInsecure(),
		grpc.WithPerRPCCredentials(preRPC{
			"authorization": "basic token",
		}),
	)
	if err != nil {
		log.Fatalln("grpc dial error:", err)
	}
	defer conn.Close()

	client := servicepb.NewServiceClient(conn)
	req := servicepb.GreetRequest{
		Name: "world",
	}
	greet, err := client.Greet(ctx, &req)
	if err != nil {
		log.Fatalln("greet error:", err)
	}
	fmt.Println("greet:", greet.Greet, "at:", greet.At.AsTime().Format(time.RFC3339))

	repReq := servicepb.RepeatGreetRequest{
		Name:         "repeat",
		TickerSecond: 1,
	}
	stream, err := client.RepeatGreet(ctx, &repReq)
	if err != nil {
		log.Fatalln("greet error:", err)
	}
	for {
		msg, err := stream.Recv()
		if err != nil {
			fmt.Println("receive error:", err)
			break
		}
		fmt.Println("receive:", msg.Greet)
	}
}
