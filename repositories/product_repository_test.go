package repositories

import (
	"ReadProducts/models"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductRepository(t *testing.T) {

	fmt.Println("Hello", db)

	productReposistory := NewProductRepository(db)

	products := []models.CreateProduct{
		{
			SourceId:      1,
			ProductTitle:  "Test-Product-1",
			ProductPrice:  100,
			StockQuantity: 10,
		},
		{
			SourceId:      2,
			ProductTitle:  "Test-Product-2",
			ProductPrice:  50,
			StockQuantity: 30,
		},
	}

	t.Run("CreateProductTest", func(t *testing.T) {

		err := productReposistory.CreateProduct(products)

		assert.Nil(t, err)

	})

}
