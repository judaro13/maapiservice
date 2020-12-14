package postcodes

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"judaro13/miaguila/models"
	"judaro13/miaguila/store"
	"judaro13/miaguila/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/avast/retry-go"
	"gorm.io/gorm"
)

func storeData(db *gorm.DB, coords [][]string, reference string) {
	chunks := splitInChuncks(coords)
	for _, chunk := range chunks {
		queryBulkData(db, chunk, reference)
	}
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

func queryBulkData(db *gorm.DB, coords [][]string, reference string) {
	query := stringCoordsToQueryStruct(coords)

	body, err := json.Marshal(query)
	if err != nil {
		utils.Error(err)
		return
	}

	APIUrl := "https://api.postcodes.io/postcodes"

	var bodyResp []byte
	var result models.UKAPIPOSTResult

	err = retry.Do(
		func() error {
			bodyIO := bytes.NewBuffer(body)
			client := http.Client{Timeout: 10 * time.Second}
			resp, err := client.Post(APIUrl, "application/json", bodyIO)
			if err != nil {
				return err
			}

			if resp.StatusCode != http.StatusOK {
				err := errors.New(string(body))
				return err
			}

			defer resp.Body.Close()
			bodyResp, err = ioutil.ReadAll(resp.Body)
			if err != nil {
				return err
			}

			err = json.Unmarshal(bodyResp, &result)
			if err != nil {
				utils.Error(err)
				return err
			}

			if result.Status == 200 && len(result.Result) == 0 {
				return errors.New("no results")
			}
			return nil
		},
		retry.Attempts(5),
		retry.Delay(5*time.Second),
	)

	if err != nil {
		utils.Error(err)
		return
	}

	store.SaveUKAPIResponse(db, result, reference)
}

func stringCoordsToQueryStruct(coords [][]string) models.UKAPIBulkQuery {
	geolocs := []models.UKAPICoordinate{}

	for _, values := range coords {
		if len(values) != 2 {
			continue
		}
		lat, err := strconv.ParseFloat(values[0], 64)
		if err != nil {
			continue
		}
		lon, err := strconv.ParseFloat(values[1], 64)
		if err != nil {
			continue
		}
		geolocs = append(geolocs, models.UKAPICoordinate{Longitude: lon, Latitude: lat, Radius: 50, Limit: 1})
	}

	return models.UKAPIBulkQuery{geolocs}
}
