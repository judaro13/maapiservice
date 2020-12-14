package postcodes

import (
	"judaro13/miaguila/apiservice/store"
	"judaro13/miaguila/apiservice/utils"
	"net/http"
	"strconv"

	"judaro13/miaguila/apiservice/models"
)

// List all stored postcodes
func List(write http.ResponseWriter, request *http.Request) {

	query := request.URL.Query()

	page, err := strconv.Atoi(query.Get("page"))
	if err != nil {
		page = 0
	}

	context := request.Context().Value("ctx").(*models.AppContext)
	result := store.ListPostcode(context.DB, page)
	utils.JSONResponse(write, models.JSONResponse{Code: models.StatusOk, Message: result})
}
