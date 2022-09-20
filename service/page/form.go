package page

import (
	"github.com/gin-gonic/gin"
	"github.com/narek-kerobian/odoo-scraper/model"
	service "github.com/narek-kerobian/odoo-scraper/service/common"
	"gorm.io/gorm"
)

// Page form struct
type PageForm struct {
    Language    string  `form:"language"`
    Title       string  `form:"title"`
    Category    string  `form:"category"`
    Subcategory string  `form:"subcategory"`
    Text        string  `form:"text"`
}

// Bind page form to gin context
func (form *PageForm) Bind(c *gin.Context) (error) {
    return c.ShouldBind(&form)
}

// Process page form
func (form PageForm) PostPageFormProcessor(pageEntity model.Page, db *gorm.DB) {
    // Return if submitted data is empty
    if form.Language == "" {
        return
    }

    // Update title
    for idx, t := range pageEntity.Title {
        if t.Lang == form.Language {
            pageEntity.Title = service.DeleteLocalizedByIndex(
                pageEntity.Title, 
                idx,
            )
        }
    }
    pageEntity.Title = append(pageEntity.Title, model.Localized{
        Lang: form.Language,
        Text: form.Title,
    })

    // Update category
    for idx, t := range pageEntity.Category {
        if t.Lang == form.Language {
            pageEntity.Category = service.DeleteLocalizedByIndex(
                pageEntity.Category,
                idx,
            )
        }
    }
    pageEntity.Category = append(pageEntity.Category, model.Localized{
        Lang: form.Language,
        Text: form.Category,
    })

    // Update subcategory
    for idx, t := range pageEntity.Subcategory {
        if t.Lang == form.Language {
            pageEntity.Subcategory = service.DeleteLocalizedByIndex(
                pageEntity.Subcategory,
                idx,
            )
        }
    }
    pageEntity.Subcategory = append(pageEntity.Subcategory, model.Localized{
        Lang: form.Language,
        Text: form.Subcategory,
    })

    // Update text
    for idx, t := range pageEntity.Text {
        if t.Lang == form.Language {
            pageEntity.Text = service.DeleteLocalizedByIndex(
                pageEntity.Text,
                idx,
            )
        }
    }
    pageEntity.Text = append(pageEntity.Text, model.Localized{
        Lang: form.Language,
        Text: form.Text,
    })

    // Persist the entity
    pageEntity.Update(db)

}
