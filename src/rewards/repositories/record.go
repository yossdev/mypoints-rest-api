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
	Category     string           `gorm:"not null"`
	Transactions []_t.Transaction `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	CreatedAt    time.Time        `gorm:"not null; default: now()"`
	UpdatedAt    time.Time        `gorm:"not null; default: now()"`
	DeletedAt    gorm.DeletedAt   `gorm:"index"`
}

type RewardCategory struct {
	Value       string   `gorm:"primaryKey; unique"`
	Description string   `gorm:"not null"`
	Rewards     []Reward `gorm:"foreignKey:Category; constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
}

func (rec *Reward) ToDomain() entities.Domain {
	return entities.Domain{
		ID:       rec.ID,
		AdminID:  rec.AdminID,
		Title:    rec.Title,
		Value:    rec.Value,
		Points:   rec.Points,
		Category: rec.Category,
	}
}

func createReward(p entities.Domain, rec *Reward) {
	rec.AdminID = p.AdminID
	rec.Title = p.Title
	rec.Value = p.Value
	rec.Points = p.Points
	rec.Category = p.Category
	if p.Img != "" {
		rec.Img = p.Img
	}
}

func updateReward(p entities.Domain, rec *Reward) {
	rec.Title = p.Title
	rec.Value = p.Value
	rec.Points = p.Points
	rec.Category = p.Category
	if p.Img != "" {
		rec.Img = p.Img
	}
}
