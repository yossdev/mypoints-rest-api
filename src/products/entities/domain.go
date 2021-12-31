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
	Points       uint32
	Img          string
	Transactions []_transaction.Domain
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
}

type Service interface {
	CreateProduct(payload Domain) (int64, error)
	UpdateProduct(productId uint32, payload Domain) int64
	DeleteProduct(productId uint32) int64
}

type PsqlRepository interface {
	Create(payload Domain) (int64, error)
	Update(payload Domain) int64
	Delete(productId uint32) int64
}
