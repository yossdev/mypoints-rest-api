package repositories

import (
	"github.com/google/uuid"
	"time"
)

type Transaction struct {
	ID      uuid.UUID `gorm:"primaryKey; type:uuid; default:uuid_generate_v4()"`
	AgentID uuid.UUID `gorm:"type:uuid; not null"`
	//RewardID      int32
	ProductID     int32
	Type          string `gorm:"not null"`
	Title         string `gorm:"not null"`
	Points        int32  `gorm:"not null"`
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
