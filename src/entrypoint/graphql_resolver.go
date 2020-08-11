package main

import (
	"cleanarchitecture-fes/src/adaptor/graphqlgen"
	"cleanarchitecture-fes/src/domain"
	"cleanarchitecture-fes/src/usecase/feseventinteractor"
	"context"
	"log"
	"strconv"
)

type graphQLResolver struct {
	FesEeventInteractor *feseventinteractor.FesEeventInteractor
}

type queryResolver struct{ *graphQLResolver }

func (r *graphQLResolver) Query() graphqlgen.QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *graphQLResolver }

func (r *graphQLResolver) Mutation() graphqlgen.MutationResolver {
	return &mutationResolver{r}
}

func (r *queryResolver) GetFesEvent(
	ctx context.Context,
) (*graphqlgen.GetFesEventPayload, error) {
	log.Printf("GetFesEvent")
	fesEvents, err := r.FesEeventInteractor.Get()
	if err != nil {
		return nil, err
	}
	log.Printf("FesEeventInteractor.Get() :%v", fesEvents)
	getFesEvents := []*graphqlgen.FesEvent{}
	for _, fesEvent := range *fesEvents {
		getFesEvents = append(getFesEvents, &graphqlgen.FesEvent{
			ID:      strconv.Itoa(fesEvent.ID),
			Title:   fesEvent.Title,
			Speaker: fesEvent.Speaker,
		})
	}
	return &graphqlgen.GetFesEventPayload{
		FesEvents: getFesEvents,
	}, nil
}

func (r *mutationResolver) SaveFesEvent(
	ctx context.Context,
	input graphqlgen.SaveFesEevntInput,
) (bool, error) {
	log.Printf("SaveFesEvent")
	err := r.FesEeventInteractor.Save(domain.FesEvent{
		Title:   input.Title,
		Speaker: input.Speaker,
	})
	if err != nil {
		return false, err
	}
	return true, nil
}
