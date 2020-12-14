package postcodes

import (
	"errors"
	"judaro13/miaguila/models"
	"judaro13/miaguila/store"
	"judaro13/miaguila/utils"
	"net/http"
	"regexp"

	"github.com/gorilla/mux"
)

func BulkStatus(write http.ResponseWriter, request *http.Request) {

	reference := mux.Vars(request)["reference"]

	if !IsValidUUID(reference) {
		utils.JSONResponseWithStatus(write, errors.New("invalid reference"), http.StatusBadRequest)
		return
	}
	context := request.Context().Value("ctx").(*models.AppContext)
	result := store.Status(context.DB, reference)
	utils.JSONResponse(write, models.JSONResponse{Code: models.StatusOk, Message: result})
}

func IsValidUUID(uuid string) bool {
	r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
	return r.MatchString(uuid)
}
