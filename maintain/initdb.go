package main

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"

	// MySqlSQL driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Db *gorm.DB

type FesEvent struct {
	gorm.Model
	Title   string `gorm:"not null"`
	Speaker string `gorm:"not null"`
}

func main() {
	var err error
	args := os.Getenv("MYSQL_ARGS")
	if args == "" {
		panic("Please set MYSQL_ARGS(read the README.)")
	}
	Db, err = gorm.Open("mysql", args)
	if err != nil {
		panic(err)
	}
	Db.AutoMigrate(&FesEvent{})
	log.Printf("AutoMigrate Success.")
}
