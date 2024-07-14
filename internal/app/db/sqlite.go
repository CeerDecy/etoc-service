package db

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SqliteDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("./etoc.db"), &gorm.Config{})
	if err != nil {
		logrus.Error(err)
		return nil
	}
	fmt.Println("init sqlite success")
	return db
}
