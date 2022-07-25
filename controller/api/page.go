package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/bracketnco/odoo-scraper/model"
	"gorm.io/gorm"
)

// Lists available pages
func GetPageCollection(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        pages := model.Page{}.FindAll(db)        
        c.JSON(http.StatusOK, pages)
    }
}

func GetPage(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        pageId := c.Param("id")
        id, err := strconv.ParseUint(pageId, 10, 32)
        if err != nil {
            panic(err)
        }

        page := model.Page{}.FindOneById(db, uint(id))
        c.JSON(http.StatusOK, page)
    }
}

func PutPage(db *gorm.DB) gin.HandlerFunc {
    return func(ctx *gin.Context) {
        
    }
}
