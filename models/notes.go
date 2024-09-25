// package models

// import (
// 	"time"

// 	"gorm.io/gorm"
// )

// type Note struct {
// 	ID        uint           `gorm:"primaryKey"`
// 	Title     string         `gorm:"size:255;not null"`
// 	Content   string         `gorm:"type:text"`
// 	UserID    uint           `gorm:"index;not null"`
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// 	DeletedAt gorm.DeletedAt `gorm:"index"`
// }


// func NoteAsll(user *User) []Note {