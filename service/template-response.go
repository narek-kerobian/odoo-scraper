package service

import (
	"github.com/gin-gonic/gin"
)

// Build tempalte and send the html response per route
func BuildTemplateResponse(c *gin.Context, status int, templateName string, data gin.H) {
    pageTempalteFile := "templates/" + templateName

    engine, _ := c.Get("engine")
    e, _ := engine.(*gin.Engine)

    e.LoadHTMLFiles(
        "templates/layouts/layout.tmpl",
        "templates/partials/header.tmpl",
        "templates/partials/footer.tmpl",
        pageTempalteFile,
    )

    c.HTML(status, templateName, data)
}
