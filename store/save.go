package store

import (
	"github.com/judaro13/masharedmodels/models"
	"gorm.io/gorm"
)

// SaveNewBulkAction storing progress
func SaveNewBulkAction(db *gorm.DB, reference string, totalRecords int) {
	progress := models.CSVUpload{Reference: reference,
		Records: totalRecords,
		Bulks:   bulkSize(totalRecords),
		Counts:  0,
		Status:  "in progress"}

	db.Create(&progress)
}

func bulkSize(totalRecords int) int {
	quotient := totalRecords / 100
	remainder := totalRecords % 100
	if remainder > 0 {
		quotient++
	}

	return quotient
}
