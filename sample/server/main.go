package main

import (
	pb "../../proto"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"log"
)

const (
	port = ":50051"
)

type Server struct {
}

func (s *Server) ApplicationCodeRegister(ctx context.Context, in *pb.Application) (*pb.ApplicationMapping, error) {
	return &pb.ApplicationMapping{
		Application: &pb.KeyWithIntegerValue {
			Key: in.ApplicationCode,
			Value: 12345,
		},
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterApplicationRegisterServiceServer(s, &Server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
