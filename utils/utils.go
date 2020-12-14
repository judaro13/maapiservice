package utils

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"judaro13/miaguila/models"
	"log"
	"net/http"
	"os"
	"runtime"
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

// NewUUID generates a random UUID according to RFC 4122, code taked from  https://play.golang.org/p/4FkNSiUDMg
func NewUUID() (string, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}
