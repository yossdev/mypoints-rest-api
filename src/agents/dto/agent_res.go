package dto

import (
	"github.com/google/uuid"
	"github.com/yossdev/mypoints-rest-api/src/agents/entities"
	"time"
)

type Profile struct {
	ID        uuid.UUID `json:"id"`
	AdminID   uuid.UUID `json:"admin_id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Points    int32     `json:"points"`
	Img       string    `json:"img"`
	Status    bool      `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomain(d *entities.Domain) Profile {
	return Profile{
		ID:        d.ID,
		AdminID:   d.AdminID,
		Name:      d.Name,
		Email:     d.Email,
		Points:    d.Points,
		Img:       d.Img,
		Status:    d.Status,
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
	}
}

type AccountCreated struct {
	RowsAffected int64 `json:"rows_affected"`
}

type AccountUpdated struct {
	RowsAffected int64 `json:"rows_affected"`
}
