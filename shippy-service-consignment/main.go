package main

import (
	"context"
	"fmt"
	pb "github.com/CoolIceV/shippy/shippy-service-consignment/proto/consignment"
	vesselProto "github.com/CoolIceV/shippy/shippy-service-vessel/proto/vessel"
	"github.com/micro/go-micro/v2"
	"log"
	"os"
)

const defaultHost = "datastore:27017"

func main() {
	service := micro.NewService(
		micro.Name("shippy.service.consignment"),
	)
	service.Init()

	uri := os.Getenv("DB_HOST")
	if uri == "" {
		uri = defaultHost
	}

	client, err := CreateClient(context.Background(), uri, 0)
	if err != nil {
		log.Panic(err)
	}
	defer client.Disconnect(context.Background())

	consignmentCollention := client.Database("shippy").Collection("consignments")

	repository := &MongoRepository{consignmentCollention}
	vesselClient := vesselProto.NewVesselService("shippy.service.vessel", service.Client())
	h := handler{repository, vesselClient}
	pb.RegisterShippingServiceHandler(service.Server(), &h)
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
	fmt.Println("consignment")
}