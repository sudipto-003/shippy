package main

import (
	"context"
	"log"
	"os"

	pb "github.com/sudipto-003/shippy/shippy-service-consignment/proto/consignment"
	vesselProto "github.com/sudipto-003/shippy/shippy-service-vessel/proto/vessel"

	"github.com/micro/go-micro/v2"
)

const (
	defaultHost = "datastore:27017"
)

func main() {
	service := micro.NewService(
		micro.Name("shippy.service.consignment"),
	)
	service.Init()
	vesselClient := vesselProto.NewVesselService("shippy.service.vessel", service.Client())

	uri := os.Getenv("DB_HOST")
	if uri == "" {
		uri = defaultHost
	}

	client, err := CreateClient(context.Background(), uri, 0)
	if err != nil {
		log.Panic(err)
	}
	defer client.Disconnect(context.Background())

	consignmentCollection := client.Database("shippy").Collection("consignment")
	repository := &MongoRepository{consignmentCollection}

	h := &handler{repository, vesselClient}
	pb.RegisterShippingServiceHandler(service.Server(), h)

	if err := service.Run(); err != nil {
		log.Panic(err)
	}
}
