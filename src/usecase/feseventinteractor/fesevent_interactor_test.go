package feseventinteractor_test

import (
	"cleanarchitecture-fes/src/domain"
	"cleanarchitecture-fes/src/usecase/feseventinteractor"
	"errors"
	"log"
	"testing"

	"github.com/tj/assert"
)

func TestFesEventUsecase(t *testing.T) {
	fesEventRepositoryMock := NewFesEventRepositoryMock()
	inputFesEvent := domain.FesEvent{
		Title:   "FesEventTestDummyNew",
		Speaker: "FesEventSpeakerDummyNew",
	}
	fesEventRepositoryMock.On(
		"Create",
		inputFesEvent,
	).Return(nil).Once()
	uc := feseventinteractor.New(fesEventRepositoryMock)

	t.Run("SaveFesEventSuccess", func(t *testing.T) {
		err := uc.Save(inputFesEvent)
		assert.Nil(t, err)
	})
	t.Run("SaveFesEventError", func(t *testing.T) {
		fesEventRepositoryMock.On(
			"Create",
			inputFesEvent,
		).Return(errors.New("DB ERROR!!!")).Once()
		err := uc.Save(inputFesEvent)
		assert.NotNil(t, err)
	})

	t.Run("GetFesEventSuccess", func(t *testing.T) {
		fesEvents := domain.FesEvents{
			{
				ID:      1,
				Title:   "hoge",
				Speaker: "SpeakerDummy",
			},
			{
				ID:      2,
				Title:   "hoge2",
				Speaker: "SpeakerDummy2",
			},
			{
				ID:      3,
				Title:   "hoge3",
				Speaker: "SpeakerDummy3",
			},
		}
		fesEventRepositoryMock.On(
			"GetAll",
		).Return(&fesEvents, nil).Once()
		events, err := uc.Get()
		assert.Nil(t, err)
		log.Print(events)
	})

	t.Run("GetFesEventError", func(t *testing.T) {
		fesEvents := domain.FesEvents{}
		fesEventRepositoryMock.On(
			"GetAll",
		).Return(&fesEvents, errors.New("GetError!!!")).Once()
		events, err := uc.Get()
		assert.NotNil(t, err)
		log.Print(events)
	})
}
