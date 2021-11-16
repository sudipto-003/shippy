package main

import (
	"context"
	"errors"

	pb "github.com/sudipto-003/shippy/shippy-service-consignment/proto/consignment"
	vesselProto "github.com/sudipto-003/shippy/shippy-service-vessel/proto/vessel"
)

type handler struct {
	repository
	vesselClient vesselProto.VesselService
}

func (s *handler) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	vesselResponse, err := s.vesselClient.FindAvailable(ctx, &vesselProto.Specification{
		MaxWeight: req.Weight,
		Capacity:  int32(len(req.Containers)),
	})

	if vesselResponse == nil {
		return errors.New("error fetching vessels, returns nil")
	}

	if err != nil {
		return err
	}

	req.VesselId = vesselResponse.Vessel.Id

	if err := s.repository.Create(ctx, MarshalConsignment(req)); err != nil {
		return err
	}

	res.Created = true
	res.Consignment = req

	return nil
}

func (s *handler) GetConsignment(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	consignments, err := s.repository.GetAll(ctx)
	if err != nil {
		return err
	}

	res.Consignments = UnmarshalConsignmentCollection(consignments)
	return nil
}
