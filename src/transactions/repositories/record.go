package repositories

import (
	"github.com/google/uuid"
	"github.com/yossdev/mypoints-rest-api/src/transactions/entities"
	"time"
)

type Transaction struct {
	ID            uuid.UUID `gorm:"primaryKey; type:uuid; default:uuid_generate_v4()"`
	AgentID       uuid.UUID `gorm:"type:uuid; not null"`
	RewardID      uint32
	ProductID     uint32
	Type          string `gorm:"not null"`
	Title         string `gorm:"not null"`
	Points        uint32 `gorm:"not null"`
	RedeemInvoice string
	Status        string    `gorm:"not null; default:Pending"`
	CreatedAt     time.Time `gorm:"not null; default: now()"`
	UpdatedAt     time.Time `gorm:"not null; default: now()"`
}

type TransactionStatus struct {
	Value        string        `gorm:"primaryKey; unique"`
	Description  string        `gorm:"not null"`
	Transactions []Transaction `gorm:"foreignKey:Status; constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
}

type TransactionType struct {
	Value        string        `gorm:"primaryKey; unique"`
	Description  string        `gorm:"not null"`
	Transactions []Transaction `gorm:"foreignKey:Type; constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
}

func (rec *Transaction) ToTransaction() entities.Domain {
	return entities.Domain{
		ID:            rec.ID,
		AgentID:       rec.AgentID,
		RewardID:      rec.RewardID,
		ProductID:     rec.ProductID,
		Type:          rec.Type,
		Title:         rec.Title,
		Points:        rec.Points,
		RedeemInvoice: rec.RedeemInvoice,
		Status:        rec.Status,
		CreatedAt:     rec.CreatedAt,
		UpdatedAt:     rec.UpdatedAt,
	}
}
