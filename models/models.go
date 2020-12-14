package models

import "gorm.io/gorm"

const (
	//StatusOk status
	StatusOk = "OK"
	//StatusErrors status
	StatusErrors = "ERRORS"
)

// JSONResponse basic struct to API response
type JSONResponse struct {
	Code    string
	Message interface{}
}

//AppContext application context struct
type AppContext struct {
	DB *gorm.DB
}
