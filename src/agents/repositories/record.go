package repositories

import (
	"github.com/google/uuid"
	"github.com/yossdev/mypoints-rest-api/src/agents/entities"
	"time"
)

type Agents struct {
	ID        uuid.UUID `gorm:"primaryKey; type:uuid; default:uuid_generate_v4()"`
	AdminID   uuid.UUID `gorm:"type:uuid; not null"`
	Name      string    `gorm:"not null"`
	Email     string    `gorm:"unique; not null"`
	Password  string    `gorm:"not null"`
	Points    int32     `gorm:"not null; default:0"`
	Img       string
	CreatedAt time.Time `gorm:"not null; default: now()"`
	UpdatedAt time.Time `gorm:"not null; default: now()"`
}

func (rec *Agents) toDomain() *entities.Domain {
	return &entities.Domain{
		ID:        rec.ID,
		AdminID:   rec.AdminID,
		Name:      rec.Name,
		Email:     rec.Email,
		Password:  rec.Password,
		Points:    rec.Points,
		Img:       rec.Img,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}
