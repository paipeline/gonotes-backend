package models

import (
	"time"

	"gorm.io/gorm"
)

// a tree like structure
type Document struct {
	ID             uint   `gorm:"primaryKey"`
	Title          string `gorm:"not null"`
	Content        string `gorm:"type:text"`
	UserID         string `gorm:"index"`
	ParentDocument *uint  `gorm:"index"`
	IsArchived     bool   `gorm:"default:false"`
	IsPublished    bool   `gorm:"default:false"`
	CoverImage     string
	Icon           string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`

	// Relationships
	Parent   *Document  `gorm:"foreignKey:ParentDocument"`
	Children []Document `gorm:"foreignKey:ParentDocument"`
}
