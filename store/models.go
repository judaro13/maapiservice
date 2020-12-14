package store

import (
	"time"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// GeoCoordinate db struct
type GeoCoordinate struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	Lat       float64 `gorm:"primaryKey" json:"lat"`
	Lon       float64 `gorm:"primaryKey" json:"lon"`
	Postcode  string  `json:"postcode"`
}
