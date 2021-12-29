package repositories

import (
	"github.com/google/uuid"
	_t "github.com/yossdev/mypoints-rest-api/src/transactions/repositories"
	"gorm.io/gorm"
	"time"
)

type Product struct {
	ID           int32     `gorm:"primaryKey; autoIncrement"`
	AdminID      uuid.UUID `gorm:"type:uuid; not null"`
	Title        string    `gorm:"not null"`
	Points       int32     `gorm:"not null"`
	Img          string
	Transactions []_t.Transaction `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	CreatedAt    time.Time        `gorm:"not null; default: now()"`
	UpdatedAt    time.Time        `gorm:"not null; default: now()"`
	DeletedAt    gorm.DeletedAt   `gorm:"index"`
}
