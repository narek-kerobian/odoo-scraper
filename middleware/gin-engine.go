package middleware

import "github.com/gin-gonic/gin"

// Expose the gin engine to be used to load html templates per route
func ExposeGinEngine(engine *gin.Engine) gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Set("engine", engine)
        c.Next()
    }
}
