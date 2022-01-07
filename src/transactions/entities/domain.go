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
	RedeemDesc       string
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

type InvoiceCallback struct {
	ID           string
	Status       string
	MerchantName string
	Amount       int
}

type Service interface {
	Claims(payload Domain) (int64, error)
	ClaimsStatus(id uuid.UUID, status string) (int64, error)
	Redeem(payload Domain) (int64, error)
	CallbackXendit(token string, payload InvoiceCallback) error
}

type PsqlRepository interface {
	GetTransaction(id string) (Domain, error)
	CreateClaims(payload Domain) (int64, error)
	UpdateClaimsStatus(id uuid.UUID, status string) (int64, error)
	CreateRedeem(payload Domain) (int64, error)
	UpdateRedeemStatus(id, status string) (Domain, error)
}
