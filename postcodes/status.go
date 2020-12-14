package postcodes

import (
	"errors"
	"judaro13/miaguila/apiservice/models"
	"judaro13/miaguila/apiservice/store"
	"judaro13/miaguila/apiservice/utils"
	"net/http"

	"github.com/gorilla/mux"
)

// BulkStatus query bulk status by reference
func BulkStatus(write http.ResponseWriter, request *http.Request) {
	reference := mux.Vars(request)["reference"]

	if !utils.IsValidUUID(reference) {
		utils.JSONResponseWithStatus(write, errors.New("invalid reference"), http.StatusBadRequest)
		return
	}
	context := request.Context().Value("ctx").(*models.AppContext)
	result := store.Status(context.DB, reference)
	utils.JSONResponse(write, models.JSONResponse{Code: models.StatusOk, Message: result})
}
