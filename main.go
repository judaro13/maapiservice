package main

import (
	"fmt"
	"judaro13/miaguila/models"
	"judaro13/miaguila/router"
	"net/http"
)

func main() {
	// db := store.ConnectToDB()
	// defer db.Close()

	context := models.AppContext{DB: nil}
	router := router.NewRouter(&context)

	// server port
	fmt.Println("Listening at port 3000")
	http.ListenAndServe(":3000", router)
}
