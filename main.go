package main

import (
	"fmt"
	"judaro13/miaguila/models"
	"judaro13/miaguila/router"
	"judaro13/miaguila/store"
	"net/http"
)

func main() {
	db := store.ConnectToDB()

	context := models.AppContext{DB: db}
	router := router.NewRouter(&context)

	// server port
	fmt.Println("Listening at port 3000")
	http.ListenAndServe(":3000", router)
}
