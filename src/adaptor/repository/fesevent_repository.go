// package repository

// import (
// 	"cleanarchitecture-fes/src/domain"

// 	"github.com/jinzhu/gorm"
// 	_ "github.com/jinzhu/gorm/dialects/postgres" // gormでpostgresを使用するために必要
// 	"golang.org/x/xerrors"
// )

// type FesEvent struct {
// 	Title   string `gorm:"not null"`
// 	Speaker string `gorm:"not null"`
// }

// type FesEventepository struct {
// 	db *gorm.DB
// }

// func New(databaseArgs string) (*FesEventepository, error) {
// 	db, err := gorm.Open(
// 		"mysql",
// 		databaseArgs,
// 	)
// 	if err != nil {
// 		return nil, xerrors.Errorf("failed to open DataBase: %v , %+v", databaseArgs, err)
// 	}
// 	return &FesEventepository{
// 		db: db,
// 	}, nil
// }

// func (r *FesEventepository) Create(fesEvemt domain.FesEvent) error {
// 	if err := r.db.Save(&fesEvemt).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (r *FesEventepository) GetAll() (*domain.FesEvents, error) {
// 	// if err := r.db.Save(&fesEvemt).Error; err != nil {
// 	// 	return nil, err
// 	// }
// 	return nil, nil
// }
