package store

import (
	"github.com/judaro13/masharedmodels/models"
	"gorm.io/gorm"
)

// ListPostcode func for return all stored geocoordinates, pagination based on ID  index
func ListPostcode(db *gorm.DB, page int) []models.GeoCoordinate {
	cooordinates := []models.GeoCoordinate{}
	init := 100 * page
	db.Where("id BETWEEN ? AND ?", init, init+100).Find(&cooordinates)
	return cooordinates
}
