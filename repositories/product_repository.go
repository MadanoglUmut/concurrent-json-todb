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

	sourceIds := make([]int, len(cProducts))

	for i, p := range cProducts {
		sourceIds[i] = p.SourceId
	}

	var existingProducts []models.Product

	if err := r.db.Table("products").Where("sourceid IN ?", sourceIds).Find(&existingProducts).Error; err != nil {
		return err
	}

	existingMap := make(map[int]models.Product)

	for _, p := range existingProducts {

		existingMap[p.SourceId] = p

	}

	var toInsert []models.CreateProduct

	for _, p := range cProducts {

		if existing, ok := existingMap[p.SourceId]; ok {

			if existing.ProductTitle != p.ProductTitle || existing.ProductPrice != p.ProductPrice || existing.StockQuantity != p.StockQuantity {
				toInsert = append(toInsert, p)
			}

		} else {

			toInsert = append(toInsert, p)

		}

	}

	if len(toInsert) > 0 {

		if err := r.db.Table("products").
			Clauses(clause.OnConflict{
				Columns: []clause.Column{{Name: "sourceid"}},
				DoUpdates: clause.Assignments(map[string]interface{}{
					"title": gorm.Expr("EXCLUDED.title"),
					"price": gorm.Expr("EXCLUDED.price"),
					"stock": gorm.Expr("EXCLUDED.stock"),
				}),
			}).
			Debug().Create(&toInsert).Error; err != nil {
			return err
		}

	}

	return nil
}

//Tek bir listede tutup clauses e atalım
