package page

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gitlab.com/bracketnco/odoo-scraper/model"
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
func (form PageForm) Bind(c *gin.Context) (error) {
    return c.ShouldBind(&form)
}

// Process page form
func (form PageForm) PostPageFormProcessor(pageEntity model.Page, db *gorm.DB) {
    fmt.Println(form)
}
