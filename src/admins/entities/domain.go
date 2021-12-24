package entities

import (
	"github.com/google/uuid"
	"github.com/yossdev/mypoints-rest-api/internal/utils/auth"
	_agent "github.com/yossdev/mypoints-rest-api/src/agents/entities"
	"time"
)

type Domain struct {
	ID        uuid.UUID
	Name      string
	Email     string
	Password  string
	Img       string
	Agents    []_agent.Domain
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Service interface {
	SignIn(payload *Domain) (auth.Token, error) // return jwt token
}

type PsqlRepository interface {
	SignInWithEmail(email string) (*Domain, error)
}
