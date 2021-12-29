package services

import "github.com/yossdev/mypoints-rest-api/src/products/entities"

type productService struct {
	productRepository entities.PsqlRepository
}

func NewProductService(p entities.PsqlRepository) entities.Service {
	return &productService{
		productRepository: p,
	}
}

func (s *productService) CreateProduct() error {
	return nil
}
