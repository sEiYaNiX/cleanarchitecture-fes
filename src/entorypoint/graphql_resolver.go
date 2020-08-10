package main

import (
	"cleanarchitecture-fes/src/adaptor/graphqlgen"
	"context"
	"log"
)

type graphQLResolver struct {
}

type queryResolver struct{ *graphQLResolver }

func (r *graphQLResolver) Query() graphqlgen.QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *graphQLResolver }

func (r *graphQLResolver) Mutation() graphqlgen.MutationResolver {
	return &mutationResolver{r}
}

func (q *queryResolver) GetFesEvent(
	ctx context.Context,
	input graphqlgen.GetFesEventInput,
) (*graphqlgen.GetFesEventPayload, error) {
	log.Printf("GetFesEvent %v:", input.ID)
	// TODO: call to Usecase
	return &graphqlgen.GetFesEventPayload{
		FesEvents: []*graphqlgen.FesEvent{
			{
				ID:       "dummy001",
				Title:    "titleDummu001",
				Speakers: "speakers001",
			},
		},
	}, nil
}

func (m *mutationResolver) SaveFesEvent(
	ctx context.Context,
	input graphqlgen.SaveFesEevntInput,
) (bool, error) {
	// TODO: call to Usecase
	return true, nil
}
