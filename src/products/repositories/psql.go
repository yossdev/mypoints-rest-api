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

func (p *productPsqlRepository) Create() error {
	return nil
}
