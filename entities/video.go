package entities

import "time"

type Video struct {
	ID          uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Title       string    `json:"title" binding:"required" gorm:"type:varchar(200)"`
	Description string    `json:"description" binding:"required,min=2,max=15" gorm:"type:varchar(200)"`
	URL         string    `json:"URL" binding:"required,url" gorm:"type:varchar(200)"`
	CreatedAt   time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
