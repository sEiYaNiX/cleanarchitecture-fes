package feseventinteractor

import (
	"cleanarchitecture-fes/src/domain"
	"log"
)

type FesEeventInteractor struct {
	FesEventRepository FesEventRepository
}

func New(
	fesEventRepository FesEventRepository,
) *FesEeventInteractor {
	return &FesEeventInteractor{
		FesEventRepository: fesEventRepository,
	}
}

func (interactor *FesEeventInteractor) Save(fesEvent domain.FesEvent) (domain.FesEvent, error) {
	// identifier, err := interactor.UserRepository.Store(u)
	// if err != nil {
	// 	return
	// }
	// user, err = interactor.UserRepository.FindById(identifier)
	log.Print("SaveCalled")
	return domain.FesEvent{}, error
}
