package jsonread

import (
	"ReadProducts/models"
	"bufio"
	"encoding/json"
	"os"
)

func LoadProductsFromFile(filePath string, batchSize int, productChan chan<- []models.CreateProduct) error {

	file, err := os.Open(filePath)

	if err != nil {
		return err
	}
	defer file.Close()

	const bufferSize = 524288

	reader := bufio.NewReaderSize(file, bufferSize)

	decoder := json.NewDecoder(reader)

	if _, err := decoder.Token(); err != nil {
		return err
	}

	var batch []models.CreateProduct

	for decoder.More() {

		var product models.CreateProduct

		if err := decoder.Decode(&product); err != nil {

			return err

		}

		batch = append(batch, product)

		if len(batch) >= batchSize {
			productChan <- batch
			batch = nil
		}

	}

	if len(batch) > 0 {
		productChan <- batch
	}

	return nil

}
