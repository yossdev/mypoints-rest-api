package repositories

import (
	"github.com/google/uuid"
	"github.com/yossdev/mypoints-rest-api/src/products/entities"
	_t "github.com/yossdev/mypoints-rest-api/src/transactions/repositories"
	"gorm.io/gorm"
	"time"
)

type Product struct {
	ID           uint32    `gorm:"primaryKey; autoIncrement"`
	AdminID      uuid.UUID `gorm:"type:uuid; not null"`
	Title        string    `gorm:"not null"`
	Points       uint32    `gorm:"not null"`
	Img          string
	Transactions []_t.Transaction `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	CreatedAt    time.Time        `gorm:"not null; default: now()"`
	UpdatedAt    time.Time        `gorm:"not null; default: now()"`
	DeletedAt    gorm.DeletedAt   `gorm:"index"`
}

func createProduct(p entities.Domain, rec *Product) {
	rec.AdminID = p.AdminID
	rec.Title = p.Title
	rec.Points = p.Points
	if p.Img != "" {
		rec.Img = p.Img
	}
}

func updateProduct(p entities.Domain, rec *Product) {
	rec.Title = p.Title
	rec.Points = p.Points
	if p.Img != "" {
		rec.Img = p.Img
	}
}
