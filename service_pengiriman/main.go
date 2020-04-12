package main

import (
	"context"
	"log"
	"net"

	pb "github.com/fahruroze/otentik/service_pengiriman/proto/pengiriman"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type IRepository interface {
	Create(*pb.Pengiriman) (*pb.Pengiriman, error)
}

type Repository struct {
	pengiriman2 []*pb.Pengiriman
}

func (repo *Repository) Create(pengiriman *pb.Pengiriman) (*pb.Pengiriman, error) {
	updated := append(repo.pengiriman2, pengiriman)
	repo.pengiriman2 = updated
	return pengiriman, nil
}

type service struct {
	repo IRepository
}

func (s *service) CreatePengiriman(ctx context.Context, req *pb.Pengiriman) (*pb.Response, error) {
	pengiriman, err := s.repo.Create(req)
	if err != nil {
		return nil, err
	}
	return &pb.Response{Created: true, Pengiriman: pengiriman}, nil
}

func main() {
	repo := &Repository{}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to connect", err)
	}
	s := grpc.NewServer()

	pb.RegisterServicePengirimanServer(s, &service{repo})

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
