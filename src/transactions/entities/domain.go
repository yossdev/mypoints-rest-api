package entities

import (
	"github.com/google/uuid"
	"time"
)

type Domain struct {
	ID            uuid.UUID
	AgentID       uuid.UUID
	RewardID      int32
	ProductID     int32
	Type          string
	Title         string
	Points        int32
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
	GetTransactions() error
	GetTransactionDetail() error
	GetTransactionsAdmin() error
}

type PsqlRepository interface {
	GetMany() error
	GetOne() error
	GetManyAdmin() error
}
