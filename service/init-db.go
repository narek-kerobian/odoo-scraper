package service

import (
	"gitlab.com/bracketnco/odoo-scraper/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Initializes the database
func InitDb(DbFilePath string) *gorm.DB {
    // Connect to the database
    db, err := gorm.Open(sqlite.Open("data/odoo_docs.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    // Audo-migrate schemas
    db.AutoMigrate(&model.Page{})

    return db
}

