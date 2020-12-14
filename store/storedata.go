package store

import (
	"judaro13/miaguila/models"

	"gorm.io/gorm"
)

// SaveUKAPIResponse
func SaveUKAPIResponse(db *gorm.DB, data models.UKAPIPOSTResult) {
	coordinates := []GeoCoordinate{}
	for _, results := range data.Result {
		for _, value := range results.Result {
			coordinates = append(coordinates, GeoCoordinate{Postcode: value.Postcode,
				Lat: value.Latitude, Lon: value.Longitude})
		}
	}

	db.CreateInBatches(coordinates, 100)
}
