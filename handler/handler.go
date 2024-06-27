package handler

import (
	"github.com/jffcm/aluraflix-backend/config"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func InitializeHandler() {
	db = config.GetSQLite()
}


