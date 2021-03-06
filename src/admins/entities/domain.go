package entities

import (
	"github.com/google/uuid"
	"github.com/yossdev/mypoints-rest-api/internal/utils/auth"
	_agent "github.com/yossdev/mypoints-rest-api/src/agents/entities"
	_product "github.com/yossdev/mypoints-rest-api/src/products/entities"
	_reward "github.com/yossdev/mypoints-rest-api/src/rewards/entities"
	"time"
)

type Domain struct {
	ID        uuid.UUID
	Name      string
	Email     string
	Password  string
	Img       string
	Agents    []_agent.Domain
	Products  []_product.Domain
	Rewards   []_reward.Domain
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Service interface {
	SignIn(payload Domain) (auth.Token, error) // return jwt token
	SignUp(payload *Domain) (int64, error)
	UpdateAdmin(id uuid.UUID, payload Domain) (int64, error)
	UpdateAvatar(id uuid.UUID, payload Domain) (int64, error)
}

type PsqlRepository interface {
	GetAdminByAgentID(id uuid.UUID) (Domain, error)
	SignInWithEmail(email string) ([2]string, error)
	CreateAdmin(payload *Domain) (int64, error)
	UpdateAdmin(payload Domain) (int64, error)
	UpdateAvatar(payload Domain) (int64, error)
}
