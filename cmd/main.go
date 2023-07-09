package main

import (
	pb "github.com/Vindori/bunker-game-go/internal/pb/api"
	"github.com/Vindori/bunker-game-go/internal/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterBunkerGameServiceServer(grpcServer, service.NewBunkerService())
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("Can't start server: %v", err)
	}
}
