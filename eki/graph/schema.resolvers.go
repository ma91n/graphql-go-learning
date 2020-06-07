package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/laqiiz/graphql-go-learning/eki/graph/generated"
	"github.com/laqiiz/graphql-go-learning/eki/graph/model"
)

func (r *queryResolver) GetStation(ctx context.Context, stationCd *int) (*model.StationConn, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetAll(ctx context.Context) ([]*model.StationConn, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
