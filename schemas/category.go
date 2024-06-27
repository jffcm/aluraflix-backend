package schemas

import (
	"strings"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Title string `json:"title" gorm:"unique" `
	Color string `json:"color"`
}

// Custom table name
func (Category) TableName() string {
	return "categories"
}

func (u *Category) BeforeSave(tx *gorm.DB) (err error) {
    u.Title = strings.ToUpper(u.Title)
    return nil
}