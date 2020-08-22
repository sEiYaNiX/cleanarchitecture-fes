package feseventinteractor_test

import (
	"cleanarchitecture-fes/src/adaptor/repository"
	"cleanarchitecture-fes/src/domain"
	"cleanarchitecture-fes/src/usecase/feseventinteractor"
	"log"
	"testing"

	"github.com/tj/assert"
)

func TestFesEventUsecase(t *testing.T) {
	args := "cafes:cafes99@(localhost:3306)/cleanarchitecture_fes"
	fesEventRepository, err := repository.New(args)
	assert.Nil(t, err)
	inputFesEvent := domain.FesEvent{
		Title:   "FesEventTestDummyNew",
		Speaker: "FesEventSpeakerDummyNew",
	}
	uc := feseventinteractor.New(fesEventRepository)
	t.Run("SaveFesEventSuccess", func(t *testing.T) {
		err = uc.Save(inputFesEvent)
		assert.Nil(t, err)
	})
	t.Run("GetFesEventSuccess", func(t *testing.T) {
		events, err := uc.Get()
		assert.Nil(t, err)
		log.Print(events)
	})
}
