package entities

import (
	"github.com/google/uuid"
	"time"
)

type Domain struct {
	ID               uuid.UUID
	AgentID          uuid.UUID
	RewardID         uint32
	ProductID        uint32
	Title            string
	Points           uint32
	NotaImg          string
	RedeemInvoiceID  string
	RedeemInvoiceURL string
	Type             string
	Status           string
	CreatedAt        time.Time
	UpdatedAt        time.Time
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
	Claims(payload Domain) (int64, error)
	ClaimsStatus(id uuid.UUID, status string) (int64, error)
}

type PsqlRepository interface {
	CreateClaims(payload Domain) (int64, error)
	UpdateClaimsStatus(id uuid.UUID, status string) (int64, error)
}
