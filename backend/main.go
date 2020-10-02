package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	"grpc_gateway/servicepb"
)

type service struct{}

func (*service) Greet(context context.Context, req *servicepb.GreetRequest) (*servicepb.GreetResponse, error) {
	fmt.Println("greet to:", req.Name)
	return &servicepb.GreetResponse{
		Greet: "Hello " + req.Name + "!",
		At:    timestamppb.Now(),
	}, nil
}

func (*service) RepeatGreet(req *servicepb.RepeatGreetRequest, svr servicepb.Service_RepeatGreetServer) error {
	fmt.Println("keep greeting to:", req.Name, "period:", req.TickerSecond)

	ticker := time.NewTicker(time.Duration(req.TickerSecond) * time.Second)
	defer ticker.Stop()

	resp := servicepb.GreetResponse{
		Greet: "Hello " + req.Name + "!",
	}
Loop:
	for {
		resp.At = timestamppb.Now()
		if err := svr.Send(&resp); err != nil {
			fmt.Println("send response error:", err)
			break
		}
		select {
		case <-svr.Context().Done():
			fmt.Println("send done")
			break Loop
		case <-ticker.C:
		}
	}

	return nil
}

func proxy(from, to string) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := servicepb.RegisterServiceHandlerFromEndpoint(ctx, mux, to, opts)
	if err != nil {
		return err
	}

	fmt.Println("Start proxy at", from)
	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(from, mux)
}

func auth(ctx context.Context) (context.Context, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		fmt.Println("incoming doesn't have metadata")
	} else {
		fmt.Println("metadata:", md)
	}
	token, err := grpc_auth.AuthFromMD(ctx, "basic")
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "no basic header found: %v", err)
	}
	fmt.Println("auth token:", token)
	if false {
		return nil, status.Errorf(codes.Unauthenticated, "invalid auth credentials: %v", err)
	}
	return ctx, nil
}

func main() {
	grpcAddr := ":50551"
	fmt.Println("Starting grpc server:", grpcAddr)

	lis, err := net.Listen("tcp", grpcAddr)

	if err != nil {
		log.Fatalf("Error while listening : %v", err)
	}

	s := grpc.NewServer(
		grpc.StreamInterceptor(grpc_auth.StreamServerInterceptor(auth)),
		grpc.UnaryInterceptor(grpc_auth.UnaryServerInterceptor(auth)),
	)
	servicepb.RegisterServiceServer(s, &service{})
	reflection.Register(s)

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Error while serving : %v", err)
		}
	}()

	proxy(":8080", grpcAddr)
}
