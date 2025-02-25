package idgenerator

import (
	"context"

	"github.com/thisPeyman/go-urlshortner/api"
	snowflakego "github.com/thisPeyman/snowflake-go"
)

type IDGeneratorService struct {
	api.UnimplementedIDGeneratorServiceServer
	node *snowflakego.Snowflake
}

func NewIDGeneratorService(node *snowflakego.Snowflake) *IDGeneratorService {
	return &IDGeneratorService{
		node: node,
	}
}

func (s *IDGeneratorService) GenerateID(context.Context, *api.GenerateIDRequest) (*api.GenerateIDResponse, error) {
	randomID, err := s.node.GenerateID()
	if err != nil {
		return nil, err
	}

	return &api.GenerateIDResponse{
		RandomID: randomID,
	}, nil
}
