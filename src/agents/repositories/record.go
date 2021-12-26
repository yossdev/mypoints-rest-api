package repositories

import (
	"github.com/google/uuid"
	"github.com/yossdev/mypoints-rest-api/src/agents/entities"
	"time"
)

type Agent struct {
	ID        uuid.UUID `gorm:"primaryKey; type:uuid; default:uuid_generate_v4()"`
	AdminID   uuid.UUID `gorm:"type:uuid; not null"`
	Name      string    `gorm:"not null"`
	Email     string    `gorm:"unique; not null"`
	Password  string    `gorm:"not null"`
	Points    int32     `gorm:"not null; default:0"`
	Img       string
	Status    bool      `gorm:"not null; default:true"`
	CreatedAt time.Time `gorm:"not null; default: now()"`
	UpdatedAt time.Time `gorm:"not null; default: now()"`
}

func (rec *Agent) toDomain() *entities.Domain {
	return &entities.Domain{
		ID:        rec.ID,
		AdminID:   rec.AdminID,
		Name:      rec.Name,
		Email:     rec.Email,
		Password:  rec.Password,
		Points:    rec.Points,
		Img:       rec.Img,
		Status:    rec.Status,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}

func updateAccount(p *entities.Domain, a *Agent) {
	a.Name = p.Name
	a.Email = p.Email
	a.Password = p.Password
	a.Img = p.Img
}
