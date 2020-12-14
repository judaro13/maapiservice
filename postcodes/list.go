package postcodes

import (
	"fmt"
	"judaro13/miaguila/models"
	"judaro13/miaguila/store"
	"judaro13/miaguila/utils"
	"net/http"
	"strconv"
)

// List all stored postcodes
func List(write http.ResponseWriter, request *http.Request) {

	query := request.URL.Query()

	page, err := strconv.Atoi(query.Get("page"))
	if err != nil {
		page = 0
	}

	fmt.Printf("\n %#v ", page)
	context := request.Context().Value("ctx").(*models.AppContext)
	result := store.ListPostcode(context.DB, page)
	utils.JSONResponse(write, models.JSONResponse{Code: models.StatusOk, Message: result})
}
