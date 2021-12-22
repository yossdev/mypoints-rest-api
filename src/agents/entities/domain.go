package entities

import (
	"github.com/google/uuid"
	"github.com/yossdev/mypoints-rest-api/internal/utils/auth"
	"time"
)

type Domain struct {
	ID        uuid.UUID
	AdminID   uuid.UUID
	Name      string
	Email     string
	Password  string
	Points    int32
	Img       string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Service interface {
	SignIn(payload *Domain) (auth.Token, error) // return jwt token
	GetAgent(id uuid.UUID) (*Domain, error)
	UpdateAgent(id uuid.UUID, payload *Domain) error
}

type PsqlRepository interface {
	SignInWithEmail(email string) (*Domain, error)
	GetAgent(id uuid.UUID) (*Domain, error)
	UpdateAgent(payload *Domain) error
}
