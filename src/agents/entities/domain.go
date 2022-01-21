package entities

import (
	"github.com/google/uuid"
	"github.com/yossdev/mypoints-rest-api/internal/utils/auth"
	_transaction "github.com/yossdev/mypoints-rest-api/src/transactions/entities"
	"time"
)

type Domain struct {
	ID           uuid.UUID
	AdminID      uuid.UUID
	Name         string
	Email        string
	Password     string
	Points       uint32
	Img          string
	Active       bool
	Transactions []_transaction.Domain
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Service interface {
	SignIn(payload Domain) (auth.Token, error) // return jwt token
	SignUp(payload *Domain) (int64, error)
	GetAgent(id uuid.UUID) (Domain, error)
	UpdateAgent(id uuid.UUID, payload Domain) (int64, error)
	UpdateAvatar(id uuid.UUID, payload Domain) (int64, error)
	UpdateAgentByAdmin(payload Domain) (int64, error)
}

type PsqlRepository interface {
	SignInWithEmail(email string) (Domain, error)
	CreateAgent(payload *Domain) (int64, error)
	GetAgent(id uuid.UUID) (Domain, error)
	UpdateAgent(payload Domain) (int64, error)
	UpdateAvatar(payload Domain) (int64, error)
	UpdatePoints(id uuid.UUID, points int32) (int64, error)
	UpdateAgentByAdmin(payload Domain) (int64, error)
}
