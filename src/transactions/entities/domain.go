package entities

import (
	"github.com/google/uuid"
	"time"
)

type Domain struct {
	ID            uuid.UUID
	AgentID       uuid.UUID
	RewardID      uint32
	ProductID     uint32
	Type          string
	Title         string
	Points        uint32
	RedeemInvoice string
	Status        string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type TransactionStatus struct {
	Value        string
	Description  string
	Transactions []Domain
}

type TransactionType struct {
	Value        string
	Description  string
	Transactions []Domain
}

type Service interface {
	CreateTransaction() error
}

type PsqlRepository interface {
	Create() error
}
