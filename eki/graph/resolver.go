package graph

import (
	"context"
	"fmt"
	"github.com/laqiiz/graphql-go-learning/eki/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{}

func (r *Resolver) getAll(ctx context.Context) ([]*model.Station, error) {

	return nil, nil
}

func (r *Resolver) getStation(ctx context.Context, stationCd *int) (*model.Station, error) {
	panic(fmt.Errorf("not implemented"))
}
