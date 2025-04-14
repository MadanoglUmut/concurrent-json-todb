package pkg

import (
	"ReadProducts/models"
	"bufio"
	"encoding/json"
	"log"
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

	log.Printf("Dosya okunuyor: %s", filePath)

	for decoder.More() {

		var product models.CreateProduct

		if err := decoder.Decode(&product); err != nil {

			return err

		}

		batch = append(batch, product)

		if len(batch) >= batchSize {
			log.Printf("Batch gönderiliyor (%s): %d ürün", filePath, len(batch))
			productChan <- batch
			batch = nil
		}

	}

	if len(batch) > 0 {
		log.Printf("Kalan batch gönderiliyor (%s): %d ürün", filePath, len(batch))
		productChan <- batch
	}

	return nil

}
