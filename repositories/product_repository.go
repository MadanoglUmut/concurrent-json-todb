package repositories

import (
	"ReadProducts/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (r *ProductRepository) CreateProduct(cProducts []models.CreateProduct) error {
	if err := r.db.Table("products").
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "sourceid"}},
			DoNothing: true,
		}).
		Create(&cProducts).Error; err != nil {
		return err
	}
	return nil
}
