package repositories

import (
	"github.com/yossdev/mypoints-rest-api/infrastuctures/db"
	"github.com/yossdev/mypoints-rest-api/src/admins/entities"
)

type adminPsqlRepository struct {
	DB db.PsqlDB
}

func NewAdminPsqlRepository(p db.PsqlDB) entities.PsqlRepository {
	return &adminPsqlRepository{
		DB: p,
	}
}

func (p *adminPsqlRepository) SignInWithEmail(email string) (*entities.Domain, error) {
	admin := Admin{}

	res := p.DB.DB().Where("email = ?", email).First(&admin)
	if res.Error != nil {
		return admin.toDomain(), res.Error
	}

	return admin.toDomain(), nil
}
