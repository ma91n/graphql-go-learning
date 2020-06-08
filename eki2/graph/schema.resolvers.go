package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/laqiiz/graphql-go-learning/eki2/graph/generated"
	"github.com/laqiiz/graphql-go-learning/eki2/graph/model"
)

func (r *queryResolver) Transfer(ctx context.Context, stationCd *int) (*model.StationConn, error) {
	return r.getStation(ctx, stationCd)
}

func (r *queryResolver) StationByName(ctx context.Context, stationName *string) ([]*model.Station, error) {
	return r.getStationByName(ctx, stationName)
}

func (r *stationConnResolver) TransferStation(ctx context.Context, obj *model.StationConn) ([]*model.Station, error) {
	return r.transferStation(ctx, obj)
}

func (r *stationConnResolver) BeforeStation(ctx context.Context, obj *model.StationConn) (*model.Station, error) {
	return r.beforeStation(ctx, obj)
}

func (r *stationConnResolver) AfterStation(ctx context.Context, obj *model.StationConn) (*model.Station, error) {
	return r.afterStation(ctx, obj)
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// StationConn returns generated.StationConnResolver implementation.
func (r *Resolver) StationConn() generated.StationConnResolver { return &stationConnResolver{r} }

type queryResolver struct{ *Resolver }
type stationConnResolver struct{ *Resolver }
