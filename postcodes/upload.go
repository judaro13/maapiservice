package postcodes

import (
	"encoding/csv"
	"io"
	"judaro13/miaguila/models"
	"judaro13/miaguila/store"
	"judaro13/miaguila/utils"
	"net/http"
)

// UploadCVS func to load csv data
func UploadCVS(write http.ResponseWriter, request *http.Request) {

	reader := csv.NewReader(request.Body)
	var results [][]string
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		results = append(results, record)
	}

	context := request.Context().Value("ctx").(*models.AppContext)

	reference, err := utils.NewUUID()
	utils.Error(err)

	go store.SaveNewBulkAction(context.DB, reference, len(results))

	go storeData(context.DB, results, reference)
	utils.JSONResponse(write, models.JSONResponse{Code: models.StatusOk,
		Message: "Processing data, Chech status with refernce " + reference})
}
