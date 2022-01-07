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

func (rec *Reward) ToDomain() entities.Domain {
	return entities.Domain{
		ID:      rec.ID,
		AdminID: rec.AdminID,
		Title:   rec.Title,
		Value:   rec.Value,
		Points:  rec.Points,
	}
}

func createReward(p entities.Domain, rec *Reward) {
	rec.AdminID = p.AdminID
	rec.Title = p.Title
	rec.Value = p.Value
	rec.Points = p.Points
	if p.Img != "" {
		rec.Img = p.Img
	}
}

func updateReward(p entities.Domain, rec *Reward) {
	rec.Title = p.Title
	rec.Value = p.Value
	rec.Points = p.Points
	if p.Img != "" {
		rec.Img = p.Img
	}
}
