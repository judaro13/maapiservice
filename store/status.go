package store

import (
	"fmt"

	"github.com/judaro13/masharedmodels/models"
	"gorm.io/gorm"
)

// Status retrieve current bulk status based on a reference
func Status(db *gorm.DB, reference string) models.CSVUpload {
	progress := models.CSVUpload{}
	fmt.Printf(" %#v ", reference)
	db.Where("reference = ?", reference).First(&progress)
	return progress
}
