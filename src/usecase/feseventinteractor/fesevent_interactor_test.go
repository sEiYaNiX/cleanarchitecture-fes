package feseventinteractor_test

import (
	"cleanarchitecture-fes/src/domain"
	"cleanarchitecture-fes/src/usecase/feseventinteractor"
	"log"
	"os"
	"testing"

	"github.com/tj/assert"
)

func TestFesEventUsecase(t *testing.T) {
	os.Setenv("MYSQL_ARGS", "cafes:cafes99@(localhost:3306)/cleanarchitecture_fes")
	t.Run("SaveFesEvent", func(t *testing.T) {
		inputFesEvent := domain.FesEvent{
			Title:   "FesEventTestDummy",
			Speaker: "FesEventSpeakerDummy",
		}
		uc := feseventinteractor.New()
		err := uc.Save(inputFesEvent)
		assert.Nil(t, err)
	})
	t.Run("GetFesEvent", func(t *testing.T) {
		uc := feseventinteractor.New()
		events, err := uc.Get()
		assert.Nil(t, err)
		log.Print(events)
	})
}
