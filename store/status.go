package store

import (
	"fmt"

	"gorm.io/gorm"
)

func Status(db *gorm.DB, reference string) CSVUpload {
	progress := CSVUpload{}
	fmt.Printf(" %#v ", reference)
	db.Where("reference = ?", reference).First(&progress)
	return progress
}
