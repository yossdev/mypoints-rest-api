package repositories

import (
	"github.com/google/uuid"
	"github.com/yossdev/mypoints-rest-api/src/admins/entities"
	_agent "github.com/yossdev/mypoints-rest-api/src/agents/entities"
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
	agents := agentSlice(rec.Agents)
	return entities.Domain{
		ID:        rec.ID,
		Name:      rec.Name,
		Email:     rec.Email,
		Password:  rec.Password,
		Agents:    agents,
		Img:       rec.Img,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}

func agentSlice(a []_a.Agent) []_agent.Domain {
	var res []_agent.Domain

	for _, val := range a {
		res = append(res, val.ToDomain())
	}
	return res
}

func createAccount(p *entities.Domain, rec *Admin) {
	rec.Name = p.Name
	rec.Email = p.Email
	rec.Password = p.Password
	rec.Img = p.Img
}

func updateAccount(p entities.Domain, rec *Admin) {
	rec.Name = p.Name
	rec.Email = p.Email
	rec.Password = p.Password
}
