package gormx

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	Id        int            `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"createdAt" gorm:"index"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"index"`
	DeletedAt gorm.DeletedAt `json:"deletedAt,omitempty" gorm:"index"`
}

type RelateModel struct {
	Id        int       `json:"id" gorm:"uniqueIndex;primaryKey"`
	CreatedAt time.Time `json:"createdAt" gorm:"index"`
	UpdatedAt time.Time `json:"updatedAt"`
}
