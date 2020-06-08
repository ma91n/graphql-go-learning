package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/laqiiz/graphql-go-learning/eki/graph/generated"
	"github.com/laqiiz/graphql-go-learning/eki/graph/model"
)

func (r *queryResolver) Transfer(ctx context.Context, stationCd *int) (*model.StationConn, error) {
	return r.getStation(ctx, stationCd)
}

func (r *queryResolver) StationByName(ctx context.Context, stationName *string) ([]*model.Station, error) {
	return r.getStationByName(ctx, stationName)
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
