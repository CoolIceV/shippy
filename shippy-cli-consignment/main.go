package main

import (
	pd "github.com/CoolIceV/shippy/shippy-service-consignment/proto/consignment"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

)

const (
	address = "localhost:50051"
	defaultFilename = "consignment.json"
)

func parseFile(file string) (pd.Consignment)
