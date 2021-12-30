package repositories

import (
	"github.com/google/uuid"
	"github.com/yossdev/mypoints-rest-api/src/admins/entities"
	_a "github.com/yossdev/mypoints-rest-api/src/agents/repositories"
	_p "github.com/yossdev/mypoints-rest-api/src/products/repositories"
	_r "github.com/yossdev/mypoints-rest-api/src/rewards/repositories"
	"time"
)

type Admin struct {
	ID        uuid.UUID `gorm:"primaryKey; type:uuid; default:uuid_generate_v4()"`
	Name      string    `gorm:"not null"`
	Email     string    `gorm:"unique; not null"`
	Password  string    `gorm:"not null"`
	Img       string
	Agents    []_a.Agent   `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	Products  []_p.Product `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	Rewards   []_r.Reward  `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	CreatedAt time.Time    `gorm:"not null; default: now()"`
	UpdatedAt time.Time    `gorm:"not null; default: now()"`
}

func (rec *Admin) ToDomain() entities.Domain {
	return entities.Domain{
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
