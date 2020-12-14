package postcodes

import (
	"fmt"
	"judaro13/miaguila/models"
	"judaro13/miaguila/store"
	"judaro13/miaguila/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func List(write http.ResponseWriter, request *http.Request) {

	page := mux.Vars(request)["page"]
	i, err := strconv.Atoi(page)
	if err != nil {
		i = 0
	}

	fmt.Printf("\n %#v ", page)
	context := request.Context().Value("ctx").(*models.AppContext)
	result := store.ListPostcode(context.DB, i)
	utils.JSONResponse(write, models.JSONResponse{Code: models.StatusOk, Message: result})
}
