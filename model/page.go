package model

import "gorm.io/gorm"

type Page struct {
    gorm.Model
    Parent      uint        `gorm:"references:ID"`
    Title       string      `gorm:"title"` 
    Text        Texts       `gorm:"serializer:json"`
    UrlOriginal string      `gorm:"url"`
}
