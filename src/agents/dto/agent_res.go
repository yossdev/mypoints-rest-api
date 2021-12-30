package dto

import (
	"github.com/google/uuid"
	"github.com/yossdev/mypoints-rest-api/src/agents/entities"
	"time"
)

type Profile struct {
	ID      uuid.UUID `json:"id"`
	AdminID uuid.UUID `json:"admin_id"`
	Name    string    `json:"name"`
	Email   string    `json:"email"`
	Points  uint32    `json:"points"`
	Img     string    `json:"img"`
	Active  bool      `json:"active"`
	//Transactions []_transaction.Domain `json:"transactions"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomainProfile(d *entities.Domain) Profile {
	return Profile{
		ID:      d.ID,
		AdminID: d.AdminID,
		Name:    d.Name,
		Email:   d.Email,
		Points:  d.Points,
		Img:     d.Img,
		Active:  d.Active,
		//Transactions: d.Transactions,
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
	}
}

type AccountCreated struct {
	RowsAffected int64 `json:"rows_affected"`
}

func FromDomainAC(res int64) AccountCreated {
	return AccountCreated{RowsAffected: res}
}

type AccountUpdated struct {
	RowsAffected int64 `json:"rows_affected"`
}

func FromDomainAU(res int64) AccountUpdated {
	return AccountUpdated{RowsAffected: res}
}
