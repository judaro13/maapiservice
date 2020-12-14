package store

import "gorm.io/gorm"

func ListPostcode(db *gorm.DB, page int) []GeoCoordinate {
	cooordinates := []GeoCoordinate{}
	init := 100 * page
	db.Where("id BETWEEN ? AND ?", init, init+100).Find(&cooordinates)
	return cooordinates
}
