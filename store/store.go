package store

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" //for use sqlite at tests
)

// ConnectToDB func that create a DB connection
func ConnectToDB() *gorm.DB {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in Connect DB", r)
		}
	}()

	db, err := gorm.Open(os.Getenv("DATABASE_TYPE"), os.Getenv("DATABASE_URL"))

	if err != nil {
		panic(err)
	}

	automigrations(db)

	return db
}

func automigrations(db *gorm.DB) {
	// db.AutoMigrate(&dbmodels.Model{})
}
