package main

import (
	"errors"
	"fmt"
	"judaro13/miaguila/apiservice/router"
	"judaro13/miaguila/apiservice/store"
	"net/http"
	"os"
	"time"

	"judaro13/miaguila/apiservice/models"

	"github.com/streadway/amqp"
)

func main() {
	validateEnvVars()

	db := store.ConnectToDB()

	conn, err := amqp.Dial(os.Getenv("RABBIT_URL"))
	defer conn.Close()
	if err != nil {
		panic(err)
	}
	go testRabbitConnection(conn)

	context := models.AppContext{DB: db, Rabbit: conn}
	router := router.NewRouter(&context)

	// server port
	fmt.Println("Listening at port 3000")
	http.ListenAndServe(":3000", router)
}

func testRabbitConnection(conn *amqp.Connection) {
	for {
		ch, err := conn.Channel()
		if err != nil {
			ch.Close()
			panic(err)
		}
		ch.Close()
		time.Sleep(30 * time.Second)
	}
}

func validateEnvVars() {
	vars := []string{"DATABASE_URL", "RABBIT_URL", "RABBIT_QUERY_DATA_QUEUE", "RABBIT_STORE_DATA_QUEUE"}
	for _, val := range vars {
		if len(val) == 0 {
			panic(errors.New("not found " + val + " environment variable"))
		}
	}
}
