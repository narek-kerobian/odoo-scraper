package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/bracketnco/odoo-scraper/model"
	"gitlab.com/bracketnco/odoo-scraper/service"
	"gorm.io/gorm"
)

// List available pages
func ListPages(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        pages := model.Page{}.FindAll(
            db.Order("category").Order("subcategory"))

        service.BuildTemplateResponse(c, http.StatusOK, "page/list.tmpl", gin.H{
            "pages": pages,
        })
    }
}

// Display and edit page texts
func EditPage(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        pageId := c.Param("id")
        id, err := strconv.ParseUint(pageId, 10, 32)
        if err != nil {
            panic(err)
        }

        page := model.Page{}.FindOneById(db, uint(id))

        service.BuildTemplateResponse(c, http.StatusOK, "page/edit.tmpl", gin.H{
            "page": page,
        })
    }
}

