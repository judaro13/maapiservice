package utils

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/judaro13/miaguila/models"
)

// Err var to add Err message to log
var Err = log.New(os.Stderr,
	"ERROR: ",
	log.Ldate|log.Ltime)

// JSONResponse func to standireze a json response with status OK
func JSONResponse(w http.ResponseWriter, response interface{}) {
	json, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

// JSONResponseWithStatus func to standireze a json response with status
func JSONResponseWithStatus(w http.ResponseWriter, response interface{}, status int) {
	json, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(json)
}

// JSONResponseErrors func to standarize API error response
func JSONResponseErrors(err error) models.JSONResponse {
	response := models.JSONResponse{
		Code:    models.StatusErrors,
		Message: err.Error(),
	}
	return response
}

// Error func for print errors
func Error(err error) {
	pc, fn, line, _ := runtime.Caller(1)
	if err != nil {
		Err.Printf("optiapi - in %s[%s:%d] %v", runtime.FuncForPC(pc).Name(), fn, line, err)
	}
}
