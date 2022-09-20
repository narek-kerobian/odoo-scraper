package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/narek-kerobian/odoo-scraper/config"
	"github.com/narek-kerobian/odoo-scraper/service"
)

// Application entry point
func main() {
    // Load env variables
    err := godotenv.Load(".env")
    if err != nil {
        panic(err)
    }

    // Get DB_PATH
    dbPath := os.Getenv("APP_DB_PATH")

    // Parse flags
    flags := service.FlagVars{}
    flags.ParseFlags()

    // Scrape oddo documentation
    if flags.Scrape {
        service.InitScraper(dbPath)
    }

    // Run the web server
    if flags.Serve {
        r := gin.Default()
        config.InitRoutes(r, dbPath)

        port := os.Getenv("APP_PORT")
        r.Run(":" + port)
    }
}
