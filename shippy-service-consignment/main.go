package main

import (
	"context"
	pd "github.com/CoolIceV/shippy/shippy-service-consignment/proto/consignment"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"sync"
)
const port = ":50051"

type respository interface {
	Create(consignment *pd.Consignment) (*pd.Consignment, error)
}

type Responsitory struct {
	mu sync.RWMutex
	consignments []*pd.Consignment
}

func (repo *Responsitory) Create(consignment *pd.Consignment) (*pd.Consignment, error) {
	repo.mu.Lock()
	updated := append(repo.consignments, consignment)
	repo.consignments = updated
	repo.mu.Unlock()
	return consignment, nil
}

type service struct {
	repo respository
}

func (s *service) CreateConsignment(ctx context.Context, req *pd.Consignment) (*pd.Response, error) {
	consignment, err := s.repo.Create(req)
	if err != nil {
		return nil, err
	}
	return &pd.Response{Created: true, Consignment: consignment}, nil
}

func main() {
	repo := &Responsitory{}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pd.RegisterShippingServiceServer(s, &service{repo})
	reflection.Register(s)
	log.Println("Running on port:", port)
	if err:= s.Serve(lis); err !=nil {
		log.Fatalf("failed to serve: %v", err)
	}
}