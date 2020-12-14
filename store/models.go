package store

import (
	"time"
)

// GeoCoordinate db struct
type GeoCoordinate struct {
	ID        uint `gorm:"autoIncrement:true,index"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Lat       float64 `gorm:"not null" json:"lat"`
	Lon       float64 `gorm:"not null" json:"lon"`
	Postcode  string  `json:"postcode"`
}
