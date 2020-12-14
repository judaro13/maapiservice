package store

import (
	"judaro13/miaguila/models"
	"judaro13/miaguila/utils"

	"gorm.io/gorm"
)

// SaveUKAPIResponse save data from UKAPI
func SaveUKAPIResponse(db *gorm.DB, data models.UKAPIPOSTResult, reference string) {
	coordinates := []GeoCoordinate{}
	for _, results := range data.Result {
		for _, value := range results.Result {
			coordinates = append(coordinates, GeoCoordinate{Postcode: value.Postcode,
				Lat: value.Latitude, Lon: value.Longitude})
		}
	}

	db.CreateInBatches(coordinates, 100)
	updateStatus(db, reference)
}

func updateStatus(db *gorm.DB, reference string) {
	progress := CSVUpload{}
	db.Where("reference = ?", reference).First(&progress)
	progress.Counts++
	if progress.Counts >= progress.Bulks {
		progress.Status = "done"
	}

	result := db.Save(&progress)
	if result.Error != nil {
		utils.Error(result.Error)
	}
}

// SaveNewBulkAction storing progress
func SaveNewBulkAction(db *gorm.DB, reference string, totalRecords int) {
	progress := CSVUpload{Reference: reference,
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
