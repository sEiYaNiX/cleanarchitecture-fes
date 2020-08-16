package feseventinteractor_test

import (
	"cleanarchitecture-fes/src/domain"
	"cleanarchitecture-fes/src/usecase/feseventinteractor"
	"errors"
	"testing"

	"github.com/tj/assert"
)

func TestCASaveFesEventUsecase(t *testing.T) {
	t.Run("FesEventを渡すと、リポジトリのSaveが正しく呼ばれる", func(t *testing.T) {
		fesEventRepositoryMock := NewFesEventRepositoryMock()
		fesEvent := domain.FesEvent{
			Title:   "TitleDummy001",
			Speaker: "FesEvents",
		}
		fesEventRepositoryMock.On(
			"Create",
			fesEvent,
		).Return(nil)
		uc := feseventinteractor.New(fesEventRepositoryMock)
		err := uc.Save(fesEvent)
		assert.Nil(t, err)
	})
	t.Run("リポジトリのSaveが失敗したらエラーを返す", func(t *testing.T) {
		fesEventRepositoryMock := NewFesEventRepositoryMock()
		fesEvent := domain.FesEvent{
			Title:   "TitleDummy001error",
			Speaker: "FesEventserror",
		}
		fesEventRepositoryMock.On(
			"Create",
			fesEvent,
		).Return(errors.New("RepoError!!!"))
		uc := feseventinteractor.New(fesEventRepositoryMock)
		err := uc.Save(fesEvent)
		assert.EqualError(t, err, "RepoError!!!")
	})
}

func TestCAGetFesEventUsecase(t *testing.T) {
	t.Run("FesEventsが返る", func(t *testing.T) {
		fesEvents := &domain.FesEvents{
			{
				ID:      1,
				Title:   "FesEvents001",
				Speaker: "Hanako",
			},
			{
				ID:      2,
				Title:   "FesEvents002",
				Speaker: "Taro",
			},
			{
				ID:      3,
				Title:   "FesEvent003",
				Speaker: "Jiro",
			},
		}
		fesEventRepositoryMock := NewFesEventRepositoryMock()
		fesEventRepositoryMock.On(
			"GetAll",
		).Return(fesEvents, nil)
		uc := feseventinteractor.New(fesEventRepositoryMock)
		retFesEvents, err := uc.Get()
		assert.Nil(t, err)
		assert.Equal(t, fesEvents, retFesEvents)
	})
}
