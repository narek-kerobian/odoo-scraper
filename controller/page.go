package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/bracketnco/odoo-scraper/model"
	service "gitlab.com/bracketnco/odoo-scraper/service/common"
	"gitlab.com/bracketnco/odoo-scraper/service/page"
	"gorm.io/gorm"
)

// List available pages
func ListPages(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        pages := model.Page{}.FindAll(
            db.Order("category").Order("subcategory"))

        // Build and return a response
        service.BuildTemplateResponse(c, http.StatusOK, "page/list.tmpl", gin.H{
            "pages": pages,
        })
    }
}

// Display and edit page texts
func EditPage(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var errors []error

        // Get page param
        pageId := c.Param("id")
        id, err := strconv.ParseUint(pageId, 10, 32)
        if err != nil {
            errors = append(errors, err)
        }

        // Retrieve page entity
        pageEntity := model.Page{}.FindOneById(db, uint(id))

        // Bind and process the submitted form
        var form page.PageForm
        if err := form.Bind(c); err != nil {
            errors = append(errors, err)
        }
        form.PostPageFormProcessor(pageEntity, db)

        // Build and return a response
        service.BuildTemplateResponse(c, http.StatusOK, "page/edit.tmpl", gin.H{
            "page": pageEntity,
            "errors": errors,
        })
    }
}

