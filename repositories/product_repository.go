package repositories

import (
	"ReadProducts/models"
	"ReadProducts/pkg/memstats"
	"ReadProducts/pkg/producthash"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProductRepository struct {
	db                  *gorm.DB
	existingProductsMap map[int]uint32
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	repo := &ProductRepository{
		db:                  db,
		existingProductsMap: make(map[int]uint32),
	}

	repo.loadAllProducts()

	return repo

}

func (r *ProductRepository) loadAllProducts() error {

	memstats.HeapStats()

	var existingProducts []models.CreateProduct

	if err := r.db.Table("products").Select("sourceid,title,price,stock").Find(&existingProducts).Error; err != nil {
		return err
	}

	for p := range existingProducts {
		r.existingProductsMap[existingProducts[p].SourceId] = producthash.HashProduct(
			existingProducts[p].ProductTitle,
			existingProducts[p].ProductPrice,
			existingProducts[p].StockQuantity,
		)
	}

	memstats.HeapStats()

	return nil

}

func (r *ProductRepository) CreateProduct(cProducts []models.CreateProduct) error {

	var toUpsert []models.CreateProduct

	for _, p := range cProducts {

		if existing, ok := r.existingProductsMap[p.SourceId]; ok {

			if existing != producthash.HashProduct(p.ProductTitle, p.ProductPrice, p.StockQuantity) {
				toUpsert = append(toUpsert, p)
			}

		} else {

			toUpsert = append(toUpsert, p)

		}

	}

	if len(toUpsert) > 0 {

		if err := r.db.Table("products").
			Clauses(clause.OnConflict{
				Columns: []clause.Column{{Name: "sourceid"}},
				DoUpdates: clause.Assignments(map[string]interface{}{
					"title": gorm.Expr("EXCLUDED.title"),
					"price": gorm.Expr("EXCLUDED.price"),
					"stock": gorm.Expr("EXCLUDED.stock"),
				}),
			}).
			Create(&toUpsert).Error; err != nil {
			return err
		}

	}

	return nil
}
