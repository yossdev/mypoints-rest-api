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

func (p *adminPsqlRepository) CreateAdmin(payload *entities.Domain) (int64, error) {
	admin := Admin{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: payload.Password,
		Img:      payload.Img,
	}
	res := p.DB.DB().Create(&admin)
	if res.Error != nil {
		return 0, res.Error
	}

	return res.RowsAffected, nil
}

func (p *adminPsqlRepository) UpdateAdmin(payload *entities.Domain) (int64, error) {
	admin := Admin{}
	p.DB.DB().First(&admin, "id = ?", payload.ID)

	updateAccount(payload, &admin)

	res := p.DB.DB().Save(&admin)
	if res.Error != nil {
		return 0, res.Error
	}

	return res.RowsAffected, nil
}

func (p *adminPsqlRepository) UpdateAvatar(payload *entities.Domain) (int64, error) {
	admin := Admin{}
	p.DB.DB().First(&admin, "id = ?", payload.ID)

	admin.Img = payload.Img

	res := p.DB.DB().Save(&admin)
	if res.Error != nil {
		return 0, res.Error
	}

	return res.RowsAffected, nil
}
