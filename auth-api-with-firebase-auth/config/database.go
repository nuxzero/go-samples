package config

import (
	"gofirebase/api"

	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func CreateDatabase() *gorm.DB {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Println(err)
		panic("Failed to connect to database!")
	}

	db.AutoMigrate((&api.User{}))

	return db
}
