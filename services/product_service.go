package services

import (
	"ReadProducts/models"
)

type productRepository interface {
	CreateProduct(cProducts []models.CreateProduct) error
}

type ProductService struct {
	productRepository productRepository
}

func NewProductService(productRepository productRepository) *ProductService {
	return &ProductService{
		productRepository: productRepository,
	}
}

func (s *ProductService) İmportProducts(productChan <-chan []models.CreateProduct) error {

	for batch := range productChan {

		if err := s.productRepository.CreateProduct(batch); err != nil {
			return err
		}

	}
	return nil

}
