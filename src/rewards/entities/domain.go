package entities

import (
	"github.com/google/uuid"
	_transaction "github.com/yossdev/mypoints-rest-api/src/transactions/entities"
	"gorm.io/gorm"
	"time"
)

type Domain struct {
	ID           uint32
	AdminID      uuid.UUID
	Title        string
	Value        uint64
	Points       uint32
	Img          string
	Transactions []_transaction.Domain
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
}

type Service interface {
	CreateReward(payload Domain) (int64, error)
	UpdateReward(rewardId uint32, payload Domain) (int64, error)
	DeleteReward(rewardId uint32) (int64, error)
}

type PsqlRepository interface {
	Create(payload Domain) (int64, error)
	Update(payload Domain) (int64, error)
	Delete(id uint32) (int64, error)
}
