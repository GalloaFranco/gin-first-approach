package entities

import (
	"time"
)

// Video GORM uses any field with the name ID as the table's primary key by default
type Video struct {
	ID          uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Title       string    `json:"title" binding:"required" gorm:"type:varchar(200)"`
	Description string    `json:"description" binding:"required,min=2,max=15" gorm:"type:varchar(200)"`
	URL         string    `json:"URL" binding:"required,url" gorm:"type:varchar(200)"`
	CreatedAt   time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

/*
gorm.Model
GORM defined a gorm.Model struct, which includes fields ID, CreatedAt, UpdatedAt, DeletedAt

// gorm.Model definition
type Model struct {
	ID        uint           `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
*/
