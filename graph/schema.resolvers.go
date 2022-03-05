package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"bocasay-gql/graph/generated"
	"bocasay-gql/graph/model"
	"context"
	"fmt"
)

func (r *mutationResolver) CreateMeteo(ctx context.Context, input model.NewMeteo) (*model.Meteo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Cities(ctx context.Context) ([]*model.City, error) {
	return Cities, nil
}

func (r *queryResolver) Meteos(ctx context.Context) ([]*model.Meteo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) MeteoTypes(ctx context.Context) ([]*model.MeteoType, error) {
	return MeteoTypes, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

var (
	Cities = []*model.City{
		{"1", "Antananaivo", []*model.Meteo{}},
		{"2", "Mahajanga", []*model.Meteo{}},
		{"3", "Antsiranana", []*model.Meteo{}},
	}
	MeteoTypes = []*model.MeteoType{
		{"1", "Pluvieux"},
		{"2", "Ensoleillé"},
		{"3", "Couvert"},
	}
)
