package model

import "gorm.io/gorm"

type Page struct {
    gorm.Model
    Parent      uint        `gorm:"references:ID"`
    Title       string      `gorm:"title"` 
    Category    string      `gorm:"category"` 
    Subcategory string      `gorm:"subcategory"` 
    Text        Texts       `gorm:"serializer:json"`
    UrlOriginal string      `gorm:"url"`
}
