package repositories

import (
	"github.com/yossdev/mypoints-rest-api/infrastuctures/db"
	"github.com/yossdev/mypoints-rest-api/src/products/entities"
)

type productPsqlRepository struct {
	DB db.PsqlDB
}

func NewProductPsqlRepository(p db.PsqlDB) entities.PsqlRepository {
	return &productPsqlRepository{
		DB: p,
	}
}

func (p *productPsqlRepository) Create(payload entities.Domain) (int64, error) {
	product := Product{}
	createProduct(payload, &product)

	res := p.DB.DB().Create(&product)
	return res.RowsAffected, res.Error
}

func (p *productPsqlRepository) Update(payload entities.Domain) int64 {
	product := Product{}
	updateProduct(payload, &product)

	res := p.DB.DB().Model(&product).Where("id = ?", payload.ID).Updates(product)
	return res.RowsAffected
}

func (p *productPsqlRepository) Delete(productId uint32) int64 {
	res := p.DB.DB().Delete(&Product{}, productId)
	return res.RowsAffected
}
