package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	service "github.com/ressley/test_task_go_middle/pkg/eventBus_v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	service.UnimplementedEventServiceServer
}

func (s *server) EventBus(ctx context.Context, req *service.Event) (*service.EmptyResponse, error) {
	log.Printf("[info] Received event with data: %v", req.GetData())
	return &service.EmptyResponse{}, nil
}

func main() {
	port := fmt.Sprintf(":%d", 50051)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("[error] Failed to listen on port %s: %v", port, err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	service.RegisterEventServiceServer(s, &server{})

	log.Printf("[info] gRPC server listening on port %s", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("[error] Failed to serve gRPC server: %v", err)
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
	fmt.Println("[info] Stopping the server")
	s.Stop()
}
