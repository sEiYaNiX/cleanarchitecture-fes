package feseventinteractor_test

import (
	"cleanarchitecture-fes/src/domain"
	"cleanarchitecture-fes/src/usecase/feseventinteractor"
	"errors"
	"testing"

	"github.com/tj/assert"
)

func TestSaveFesEventUsecase(t *testing.T) {
	t.Run("FesEventを渡すとリポジトリのSaveが正しく呼ばれる", func(t *testing.T) {
		inputFesEvent := domain.FesEvent{
			Title:   "FesEventTestDummy",
			Speaker: "FesEventSpeakerDummy",
		}
		createdFesEvent := &domain.FesEvent{
			ID:      1,
			Title:   "FesEventTestDummy",
			Speaker: "FesEventSpeakerDummy",
		}
		fesEventRepositoryMock := NewFesEventRepositoryMock()
		fesEventRepositoryMock.On(
			"Create",
			inputFesEvent,
		).Return(createdFesEvent, nil)

		uc := feseventinteractor.New(fesEventRepositoryMock)
		event, err := uc.Save(inputFesEvent)
		assert.Nil(t, err)
		assert.Equal(t, createdFesEvent, event)
	})
	t.Run("リポジトリSaveが失敗したとき、リポジトリが返すエラーを得る", func(t *testing.T) {
		inputFesEvent := domain.FesEvent{
			Title:   "FesEventTestDummyError",
			Speaker: "FesEventSpeakerDummyError",
		}
		fesEventRepositoryMock := NewFesEventRepositoryMock()
		fesEventRepositoryMock.On(
			"Create",
			inputFesEvent,
		).Return(nil, errors.New("RepositoryCreateError"))
		uc := feseventinteractor.New(fesEventRepositoryMock)
		_, err := uc.Save(inputFesEvent)
		assert.EqualError(t, err, "RepositoryCreateError")
	})
}

func TestGetFesEventsUsecase(t *testing.T) {
	fesEvents := &domain.FesEvents{
		{
			ID:      1,
			Title:   "FesEventTestDummy001",
			Speaker: "FesEventSpeakerDummy001",
		},
		{
			ID:      2,
			Title:   "FesEventTestDummy002",
			Speaker: "FesEventSpeakerDummy002",
		},
		{
			ID:      3,
			Title:   "FesEventTestDummy003",
			Speaker: "FesEventSpeakerDummy003",
		},
	}
	fesEventRepositoryMock := NewFesEventRepositoryMock()
	fesEventRepositoryMock.On(
		"GetAll",
	).Return(fesEvents, nil)

	uc := feseventinteractor.New(fesEventRepositoryMock)
	events, err := uc.Get()
	assert.Nil(t, err)
	assert.Equal(t, fesEvents, events)
}
