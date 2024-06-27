package schemas

import (
	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	Title       string   `json:"title"`
	Description string   `json:"description"`
	URL         string   `json:"url"`
	CategoryID  uint     `json:"category_id" gorm:"default:1"`
	Category    Category `json:"-" gorm:"foreignkey:CategoryID"`
}

// Custom table name
func (Video) TableName() string {
	return "videos"
}
