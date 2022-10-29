package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	BaseInfo
	Name        string       `gorm:"size:255" json:"name"`
	Email       string       `gorm:"size:255" json:"email"`
	CreatedBy   string       `gorm:"size:255;default:system" json:"createdBy"`
	UpdatedBy   string       `gorm:"size:255" json:"updateBy"`
	CreditCards []CreditCard `gorm:"foreignKey:UserID" json:"creditCards"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}

type CreditCard struct {
	gorm.Model
	Number string `gorm:"size:255" json:"number"`
	UserID uuid.UUID
}
