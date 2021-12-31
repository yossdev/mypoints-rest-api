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

func (s *productService) CreateProduct(payload entities.Domain) (int64, error) {
	res, err := s.productRepository.Create(payload)
	return res, err
}

func (s *productService) UpdateProduct(productId uint32, payload entities.Domain) int64 {
	payload.ID = productId
	res := s.productRepository.Update(payload)
	return res
}

func (s *productService) DeleteProduct(productId uint32) int64 {
	res := s.productRepository.Delete(productId)
	return res
}
