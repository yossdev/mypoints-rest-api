package repositories

import (
	"github.com/google/uuid"
	"github.com/yossdev/mypoints-rest-api/src/admins/entities"
	_agent "github.com/yossdev/mypoints-rest-api/src/agents/repositories"
	"time"
)

type Admin struct {
	ID        uuid.UUID `gorm:"primaryKey; type:uuid; default:uuid_generate_v4()"`
	Name      string    `gorm:"not null"`
	Email     string    `gorm:"unique; not null"`
	Password  string    `gorm:"not null"`
	Img       string
	Agents    []_agent.Agent `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	CreatedAt time.Time      `gorm:"not null; default: now()"`
	UpdatedAt time.Time      `gorm:"not null; default: now()"`
}

func (rec *Admin) toDomain() *entities.Domain {
	return &entities.Domain{
		ID:        rec.ID,
		Name:      rec.Name,
		Email:     rec.Email,
		Password:  rec.Password,
		Img:       rec.Img,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}

func updateAccount(p *entities.Domain, a *Admin) {
	a.Name = p.Name
	a.Email = p.Email
	if p.Password != "" {
		a.Password = p.Password
	}
}
