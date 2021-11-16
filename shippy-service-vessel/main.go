package main

import (
	"context"
	"log"
	"os"

	"github.com/micro/go-micro/v2"
	pb "github.com/sudipto-003/shippy/shippy-service-vessel/proto/vessel"
)


func populateDummyVessel(repo repository) {
	vessels := []*Vessel{
		{ID: "vessel001", Name: "Booty McBoatFace", MaxWeight: 20000, Capacity: 500},
	}
	for _, v := range vessels {
		repo.Create(context.Background(), v)
	}
}

func main() {
	service := micro.NewService(micro.Name("shippy.service.vessel"))
	service.Init()

	uri := os.Getenv("DB_HOST")
	if uri == "" {
		uri = "datastore:27017"
	}

	client, err := CreateClient(context.Background(), uri, 0)
	if err != nil {
		log.Panic(err)
	}
	defer client.Disconnect(context.Background())

	vesselCollection := client.Database("shippy").Collection("vessel")
	repository := &MongoRepository{vesselCollection}

	populateDummyVessel(repository)

	h := &handler{repository}

	if err := pb.RegisterVesselServiceHandler(service.Server(), h); err != nil {
		log.Panic(err)
	}

	if err := service.Run(); err != nil {
		log.Panic(err)
	}
}
