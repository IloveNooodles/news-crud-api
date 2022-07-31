package routes

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RoutesInit(server *gin.Engine, db *sql.DB) {
	api := server.Group("/api")
	v1 := api.Group("/v1")
	v1.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hellow",
		})
	})
}
