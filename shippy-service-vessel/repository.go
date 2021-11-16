package main

import (
	"context"

	pb "github.com/sudipto-003/shippy/shippy-service-vessel/proto/vessel"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Vessel struct {
	ID        string `json:"id"`
	Capacity  int32  `json:"capacity"`
	MaxWeight int32  `json:"max_weight"`
	Name      string `json:"name"`
	Available bool   `json:"available"`
	OwnerID   string `json:"owner_id"`
}

type Specification struct {
	Capacity  int32 `json:"capacity"`
	MaxWeight int32 `json:"max_weight"`
}

func MarshalSpecification(spec *pb.Specification) *Specification {
	return &Specification{
		Capacity:  spec.Capacity,
		MaxWeight: spec.MaxWeight,
	}
}

func UnmarshalSpecification(spec *Specification) *pb.Specification {
	return &pb.Specification{
		Capacity:  spec.Capacity,
		MaxWeight: spec.MaxWeight,
	}
}

func MarshalVessel(vessel *pb.Vessel) *Vessel {
	return &Vessel{
		ID:        vessel.Id,
		Capacity:  vessel.Capacity,
		MaxWeight: vessel.MaxWeight,
		Name:      vessel.Name,
		Available: vessel.Available,
		OwnerID:   vessel.OwnerId,
	}
}

func UnmarshalVessel(vessel *Vessel) *pb.Vessel {
	return &pb.Vessel{
		Id:        vessel.ID,
		Capacity:  vessel.Capacity,
		MaxWeight: vessel.MaxWeight,
		Name:      vessel.Name,
		Available: vessel.Available,
		OwnerId:   vessel.OwnerID,
	}
}

type repository interface {
	FindAvailable(ctx context.Context, spec *Specification) (*Vessel, error)
	Create(ctx context.Context, vessel *Vessel) error
}

type MongoRepository struct {
	collection *mongo.Collection
}

func (repository *MongoRepository) FindAvailable(ctx context.Context, spec *Specification) (*Vessel, error) {
	// filter := []bson.M{{
	// 	"capacity":   []bson.M{{"$lte": spec.Capacity}},
	// 	"max_weight": []bson.M{{"$lte": spec.MaxWeight}},
	// }}
	filter := bson.D{
		{"$and",
			bson.A{
				bson.D{{"capacity", bson.D{{"$gte", spec.Capacity}}}},
				bson.D{{"maxweight", bson.D{{"$gte", spec.MaxWeight}}}},
			}},
	}

	vessel := &Vessel{}
	if err := repository.collection.FindOne(ctx, filter).Decode(vessel); err != nil {
		return nil, err
	}

	return vessel, nil
}

func (repository *MongoRepository) Create(ctx context.Context, vessel *Vessel) error {
	_, err := repository.collection.InsertOne(ctx, vessel)
	return err
}
