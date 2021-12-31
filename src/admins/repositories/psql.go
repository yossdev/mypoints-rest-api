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

func (p *adminPsqlRepository) SignInWithEmail(email string) ([2]string, error) {
	admin := Admin{}

	res := p.DB.DB().Where("email = ?", email).First(&admin)
	if res.Error != nil {
		return [2]string{"", ""}, res.Error
	}

	return [2]string{admin.Password, admin.ID.String()}, nil
}

func (p *adminPsqlRepository) CreateAdmin(payload *entities.Domain) (int64, error) {
	admin := Admin{}
	createAccount(payload, &admin)

	res := p.DB.DB().Create(&admin)
	return res.RowsAffected, res.Error
}

func (p *adminPsqlRepository) UpdateAdmin(payload entities.Domain) (int64, error) {
	admin := Admin{}
	updateAccount(payload, &admin)

	res := p.DB.DB().Model(&admin).Where("id = ?", payload.ID).Updates(admin)
	return res.RowsAffected, res.Error
}

func (p *adminPsqlRepository) UpdateAvatar(payload entities.Domain) (int64, error) {
	admin := Admin{}
	admin.Img = payload.Img

	res := p.DB.DB().Model(&admin).Where("id = ?", payload.ID).Updates(admin)
	return res.RowsAffected, res.Error
}
