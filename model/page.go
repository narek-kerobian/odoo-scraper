package model

import "gorm.io/gorm"

type Page struct {
    BaseModel
    Parent      uint          `gorm:"references:ID"   json:"parent_id"`
    Title       LocalizedList `gorm:"serializer:json" json:"title"` 
    Category    LocalizedList `gorm:"serializer:json" json:"category"` 
    Subcategory LocalizedList `gorm:"serializer:json" json:"subcategory"` 
    Text        LocalizedList `gorm:"serializer:json" json:"text"`
    UrlOriginal string        `gorm:"url"             json:"url_original"`
}

// Get single record by id
func (Page) FindOneById(db *gorm.DB, id uint) (result Page) {
    db.First(&result, id)
    return
}

// Retrieve all records
func (Page) FindAll(db *gorm.DB) (result []Page) {
    db.Find(&result)
    return
}

func (page Page) Update(db *gorm.DB) {
    db.Save(&page) 
    return
}
