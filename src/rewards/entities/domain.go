package entities

import (
	"github.com/google/uuid"
	_transaction "github.com/yossdev/mypoints-rest-api/src/transactions/entities"
	"gorm.io/gorm"
	"time"
)

type Domain struct {
	ID           int32
	AdminID      uuid.UUID
	Title        string
	Value        int64
	Points       int32
	Img          string
	Transactions []_transaction.Domain
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
}

type Service interface {
	CreateReward() error
}

type PsqlRepository interface {
	Create() error
}