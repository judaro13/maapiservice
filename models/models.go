package models

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
