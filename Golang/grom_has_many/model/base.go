package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseInfo struct {
	ID        uuid.UUID      `json:"id" gorm:"type:char(36);primary_key"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
