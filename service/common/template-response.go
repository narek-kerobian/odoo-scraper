package service

import (
	"text/template"

	"github.com/gin-gonic/gin"
)

// Build tempalte and send the html response per route
func BuildTemplateResponse(c *gin.Context, status int, templateName string, data gin.H) {
    pageTempalteFile := "templates/" + templateName

    // Get gin engine
    engine, _ := c.Get("engine")
    e, _ := engine.(*gin.Engine)

    // Load custom functions
    e.SetFuncMap(template.FuncMap{
        "ParseRawHtml": ParseRawHtml,
        "GetAppLanguages": GetAppLanguages,
    })

    // Load html templates
    e.LoadHTMLFiles(
        "templates/layouts/layout.tmpl",
        "templates/partials/errors.tmpl",
        "templates/partials/header.tmpl",
        pageTempalteFile,
    )

    // Return the response
    c.HTML(status, templateName, data)
}
