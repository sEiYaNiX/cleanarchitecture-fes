package repository

import (
	"cleanarchitecture-fes/src/domain"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"golang.org/x/xerrors"
)

type FesEvent struct {
	ID      uint   `gorm:"not null"`
	Title   string `gorm:"not null"`
	Speaker string `gorm:"not null"`
}

type FesEventepository struct {
	db *gorm.DB
}

func New(databaseArgs string) (*FesEventepository, error) {
	db, err := gorm.Open(
		"mysql",
		databaseArgs,
	)
	if err != nil {
		return nil, xerrors.Errorf("failed to open DataBase: %v , %+v", databaseArgs, err)
	}
	return &FesEventepository{
		db: db,
	}, nil
}

func (r *FesEventepository) Create(fesEvemt domain.FesEvent) error {
	if err := r.db.Save(&fesEvemt).Error; err != nil {
		return err
	}
	return nil
}

func (r *FesEventepository) GetAll() (*domain.FesEvents, error) {
	fesEvents := []FesEvent{}
	errs := r.db.Find(&fesEvents).GetErrors()
	if len(errs) > 0 {
		for _, err := range errs {
			fmt.Printf("GetALl error: %+v", err)
		}
		return nil, errs[0]
	}
	getFesEvents := domain.FesEvents{}
	for _, fesEvent := range fesEvents {
		getFesEvents = append(getFesEvents, domain.FesEvent{
			ID:      int(fesEvent.ID),
			Title:   fesEvent.Title,
			Speaker: fesEvent.Speaker,
		})
	}

	return &getFesEvents, nil
}
