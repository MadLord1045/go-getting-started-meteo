package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"bocasay-gql/graph/generated"
	"bocasay-gql/graph/model"
	"context"
	"errors"
	"strconv"
)

func (r *mutationResolver) CreateMeteo(ctx context.Context, input model.NewMeteo) (*model.Meteo, error) {
	id := input.CityID
	var meteoType int
	for index, item := range MeteoTypes {
		if item.ID == input.MeteoType {
			meteoType = index
		}
	}
	for _, item := range Cities {
		if item.ID == id {
			result := &model.Meteo{
				ID:        strconv.FormatInt(int64(len(item.Meteo)+2), 10),
				Day:       input.Day,
				MeteoType: MeteoTypes[meteoType],
			}
			item.Meteo = append(item.Meteo, result)
			return result, nil
		}
	}
	return nil, errors.New("Mismatch ids")
}

func (r *queryResolver) Cities(ctx context.Context) ([]*model.City, error) {
	return Cities, nil
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
	MeteoTypes = []*model.MeteoType{
		{"1", "Pluvieux"},
		{"2", "Ensoleill√©"},
		{"3", "Couvert"},
	}
	Cities = []*model.City{
		{"1", "Antananaivo", []*model.Meteo{}},
		{"2", "Mahajanga", []*model.Meteo{}},
		{"3", "Antsiranana", []*model.Meteo{
			{"1", "Lundi", MeteoTypes[0]},
			{"2", "Mardi", MeteoTypes[1]},
			{"3", "Mercredi", MeteoTypes[2]},
		}},
	}
)
