package repositories

import (
	"github.com/google/uuid"
	"github.com/yossdev/mypoints-rest-api/src/agents/entities"
	_transaction "github.com/yossdev/mypoints-rest-api/src/transactions/entities"
	_t "github.com/yossdev/mypoints-rest-api/src/transactions/repositories"
	"time"
)

type Agent struct {
	ID           uuid.UUID `gorm:"primaryKey; type:uuid; default:uuid_generate_v4()"`
	AdminID      uuid.UUID `gorm:"type:uuid; not null"`
	Name         string    `gorm:"not null"`
	Email        string    `gorm:"unique; not null"`
	Password     string    `gorm:"not null"`
	Points       uint32    `gorm:"not null; default:0"`
	Img          string
	Active       bool             `gorm:"not null; default:true"`
	Transactions []_t.Transaction `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	CreatedAt    time.Time        `gorm:"not null; default: now()"`
	UpdatedAt    time.Time        `gorm:"not null; default: now()"`
}

func (rec *Agent) ToDomain() entities.Domain {
	//transaction := transactionSlice(rec.Transactions)
	return entities.Domain{
		ID:       rec.ID,
		AdminID:  rec.AdminID,
		Name:     rec.Name,
		Email:    rec.Email,
		Password: rec.Password,
		Points:   rec.Points,
		Img:      rec.Img,
		Active:   rec.Active,
		//Transactions: transaction,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}

func createAccount(p *entities.Domain, rec *Agent) {
	rec.AdminID = p.AdminID
	rec.Name = p.Name
	rec.Email = p.Email
	rec.Password = p.Password
	if p.Img != "" {
		rec.Img = p.Img
	}
	rec.Active = p.Active
}

func updateAccount(p entities.Domain, rec *Agent) {
	rec.Name = p.Name
	rec.Email = p.Email
	if p.Password != "" {
		rec.Password = p.Password
	}
}

func transactionSlice(t []_t.Transaction) []_transaction.Domain {
	var res []_transaction.Domain

	for _, val := range t {
		res = append(res, val.ToTransaction())
	}
	return res
}
