package config

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/bracketnco/odoo-scraper/controller"
	"gitlab.com/bracketnco/odoo-scraper/controller/api"
	"gitlab.com/bracketnco/odoo-scraper/service"
)

var apiPrefix = "/api/v1"

// Defines application routes
func InitRoutes(r *gin.Engine, dbPath string) {
    db := service.InitDb(dbPath)

    // Heartbeat controller
    r.GET("/heartbeat", controller.GetHeartbeat)

    // Load templates
    r.LoadHTMLGlob("templates/**/*")

    // API routes
    // Page routes
    gr := r.Group(apiPrefix) 
    {
        gr.GET("/page",      api.GetPageCollection(db)) 
        gr.GET("/page/:id",  api.GetPage(db)) 
        gr.PUT("/page",      api.PutPage(db)) 
    }

    // UI routes
    // Page routes
    r.GET("",  controller.ListPages(db)) 
    r.GET("/:id", controller.EditPage(db)) 
    // r.POST("/:id", controller.EditPage(db)) 
}

