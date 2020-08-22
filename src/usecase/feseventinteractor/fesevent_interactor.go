package feseventinteractor

import (
	"cleanarchitecture-fes/src/adaptor/repository"
	"cleanarchitecture-fes/src/domain"
	"os"
)

type FesEeventInteractor struct {
}

func New() *FesEeventInteractor {
	return &FesEeventInteractor{}
}

func (interactor *FesEeventInteractor) Save(fesEvent domain.FesEvent) error {
	args := os.Getenv("MYSQL_ARGS")
	if args == "" {
		panic("Please set MYSQL_ARGS(read the README.)")
	}

	fesEventRepository, err := repository.New(args)
	if err != nil {
		return err
	}

	err = fesEventRepository.Create(fesEvent)
	if err != nil {
		return err
	}
	return nil
}

func (interactor *FesEeventInteractor) Get() (*domain.FesEvents, error) {
	args := os.Getenv("MYSQL_ARGS")
	if args == "" {
		panic("Please set MYSQL_ARGS(read the README.)")
	}

	fesEventRepository, err := repository.New(args)
	if err != nil {
		return nil, err
	}

	createdFesEvent, err := fesEventRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return createdFesEvent, nil
}
