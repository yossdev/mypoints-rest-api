package repositories

import (
	"github.com/google/uuid"
	"github.com/yossdev/mypoints-rest-api/src/transactions/entities"
	"time"
)

type Transaction struct {
	ID               uuid.UUID `gorm:"primaryKey; type:uuid; default:uuid_generate_v4()"`
	AgentID          uuid.UUID `gorm:"type:uuid; not null"`
	RewardID         uint32
	ProductID        uint32
	Title            string `gorm:"not null"`
	Points           uint32 `gorm:"not null"`
	NotaImg          string
	RedeemInvoiceID  string `gorm:"unique"`
	RedeemInvoiceURL string
	RedeemDesc       string
	Type             string    `gorm:"not null"`
	Status           string    `gorm:"not null; default:Pending"`
	CreatedAt        time.Time `gorm:"not null; default: now()"`
	UpdatedAt        time.Time `gorm:"not null; default: now()"`
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

func (rec *Transaction) ToDomain() entities.Domain {
	return entities.Domain{
		ID:               rec.ID,
		AgentID:          rec.AgentID,
		RewardID:         rec.RewardID,
		ProductID:        rec.ProductID,
		Title:            rec.Title,
		Points:           rec.Points,
		NotaImg:          rec.NotaImg,
		RedeemInvoiceID:  rec.RedeemInvoiceID,
		RedeemInvoiceURL: rec.RedeemInvoiceURL,
		Type:             rec.Type,
		Status:           rec.Status,
		CreatedAt:        rec.CreatedAt,
		UpdatedAt:        rec.UpdatedAt,
	}
}

func createClaims(p entities.Domain, rec *Transaction) {
	rec.AgentID = p.AgentID
	rec.ProductID = p.ProductID
	rec.Title = p.Title
	rec.Points = p.Points
	rec.NotaImg = p.NotaImg
	rec.Type = "Debit"
}

func createRedeem(p entities.Domain, rec *Transaction) {
	rec.AgentID = p.AgentID
	rec.RewardID = p.RewardID
	rec.Title = p.Title
	rec.Points = p.Points
	rec.RedeemInvoiceID = p.RedeemInvoiceID
	rec.RedeemInvoiceURL = p.RedeemInvoiceURL
	rec.RedeemDesc = p.RedeemDesc
	rec.Type = "Credit"
}
