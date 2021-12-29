package repositories

import (
	"github.com/google/uuid"
	"github.com/yossdev/mypoints-rest-api/src/rewards/entities"
	_t "github.com/yossdev/mypoints-rest-api/src/transactions/repositories"
	"gorm.io/gorm"
	"time"
)

type Reward struct {
	ID           uint32    `gorm:"primaryKey; autoIncrement"`
	AdminID      uuid.UUID `gorm:"type:uuid; not null"`
	Title        string    `gorm:"not null"`
	Value        uint64    `gorm:"not null"`
	Points       uint32    `gorm:"not null"`
	Img          string
	Transactions []_t.Transaction `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	CreatedAt    time.Time        `gorm:"not null; default: now()"`
	UpdatedAt    time.Time        `gorm:"not null; default: now()"`
	DeletedAt    gorm.DeletedAt   `gorm:"index"`
}

func updateReward(p *entities.Domain, r *Reward) {
	r.Title = p.Title
	r.Value = p.Value
	r.Points = p.Points
	if r.Img != "" {
		r.Img = p.Img
	}
}
