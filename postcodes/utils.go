package postcodes

import (
	"encoding/json"
	"os"

	"github.com/judaro13/masharedmodels/utils"

	"judaro13/miaguila/apiservice/models"

	smodels "github.com/judaro13/masharedmodels/models"

	"github.com/streadway/amqp"
)

func storeData(context *models.AppContext, coords [][]string, reference string) {
	chunks := splitInChuncks(coords)
	for _, chunk := range chunks {
		sendDataToProcess(context.Rabbit, chunk, reference)
	}
}

func sendDataToProcess(rabbit *amqp.Connection, coords [][]string, reference string) {
	data := smodels.QueryCoordinatesMessage{Reference: reference, Coordinates: coords}
	body, err := json.Marshal(data)
	if err != nil {
		utils.Error(err)
		return
	}

	ch, err := rabbit.Channel()
	defer ch.Close()

	q, err := ch.QueueDeclare(os.Getenv("RABBIT_PROCESS_DATA_QUEUE"), false, false, false, false, nil)
	utils.Error(err)

	err = ch.Publish("", q.Name, false, false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})

	utils.Error(err)
}

func splitInChuncks(slice [][]string) [][][]string {
	chunkSize := 100
	var chunks [][][]string
	for {
		if len(slice) == 0 {
			break
		}
		if len(slice) < chunkSize {
			chunkSize = len(slice)
		}

		chunks = append(chunks, slice[0:chunkSize])
		slice = slice[chunkSize:]
	}
	return chunks
}
