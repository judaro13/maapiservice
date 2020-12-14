package postcodes

import (
	"encoding/csv"
	"io"
	"judaro13/miaguila/apiservice/store"
	"judaro13/miaguila/apiservice/utils"
	"net/http"

	"judaro13/miaguila/apiservice/models"
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
	if err != nil {
		utils.JSONResponse(write, models.JSONResponse{Code: models.StatusErrors,
			Message: err.Error()})
		return
	}

	go store.SaveNewBulkAction(context.DB, reference, len(results))

	utils.JSONResponse(write, models.JSONResponse{Code: models.StatusOk,
		Message: "Processing data, Chech status with GET /postcodes/" + reference})
}
