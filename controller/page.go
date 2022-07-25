package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/bracketnco/odoo-scraper/model"
	"gorm.io/gorm"
)

func ListPages(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        pages := model.Page{}.FindAll(
            db.Order("category").Order("subcategory"))
        c.HTML(http.StatusOK, "page/list.tmpl", gin.H{
            "pages": pages,
		})
    }
}

func EditPage(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        pageId := c.Param("id")
        id, err := strconv.ParseUint(pageId, 10, 32)
        if err != nil {
            panic(err)
        }

        page := model.Page{}.FindOneById(db, uint(id))

        c.HTML(http.StatusOK, "page/edit.tmpl", gin.H{
            "page": page,
		})
    }
}

