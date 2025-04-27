package jsonread

import (
	"ReadProducts/models"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadProductsFromFile(t *testing.T) {
	t.Run("Başarılı Yükleme", func(t *testing.T) {
		tempFile, err := os.CreateTemp("", "test_products_*.json")
		assert.Nil(t, err)
		defer os.Remove(tempFile.Name())

		testData := `[
			{"id": 1, "product_title": "Ürün 1", "product_price": 100, "stock_quantity":10},
			{"id": 2, "product_title": "Ürün 2", "product_price": 200, "stock_quantity":20},
			{"id": 3, "product_title": "Ürün 3", "product_price": 300, "stock_quantity":30},
			{"id": 4, "product_title": "Ürün 4", "product_price": 400, "stock_quantity":40}
		]`
		_, err = tempFile.WriteString(testData)
		assert.Nil(t, err)
		tempFile.Close()

		productChan := make(chan []models.CreateProduct, 2)
		batchSize := 2

		go func() {
			err := LoadProductsFromFile(tempFile.Name(), batchSize, productChan)
			assert.Nil(t, err)
			close(productChan)
		}()

		var batches [][]models.CreateProduct
		for batch := range productChan {
			batches = append(batches, batch)
		}

		assert.Equal(t, 2, len(batches), "Batch sayısı eşleşmiyor")
		assert.Len(t, batches[0], 2, "İlk batch boyutu eşleşmiyor")
		assert.Len(t, batches[1], 2, "İkinci batch boyutu eşleşmiyor")
	})

	t.Run("Geçersiz JSON dosyası", func(t *testing.T) {
		tempFile, err := os.CreateTemp("", "test_products_error_*.json")
		assert.Nil(t, err)
		defer os.Remove(tempFile.Name())

		_, err = tempFile.WriteString("geçersiz json içeriği")
		assert.Nil(t, err)
		tempFile.Close()

		productChan := make(chan []models.CreateProduct, 1)
		err = LoadProductsFromFile(tempFile.Name(), 10, productChan)
		assert.Error(t, err)
	})

	t.Run("Olmayan dosya", func(t *testing.T) {
		productChan := make(chan []models.CreateProduct, 1)
		err := LoadProductsFromFile("olmayan_dosya.json", 10, productChan)
		assert.Error(t, err)
	})
}
