package routes

import (
	"database/sql"
	"net/http"

	"github.com/IloveNooodles/kumparan-techincal-test/internal/handler"
	"github.com/IloveNooodles/kumparan-techincal-test/internal/repository"
	"github.com/IloveNooodles/kumparan-techincal-test/internal/service"
	"github.com/gin-gonic/gin"
)

func RoutesInit(server *gin.Engine, db *sql.DB) {
	api := server.Group("/api")
	v1 := api.Group("/v1")

	authorsRepository := repository.NewAuthorsRepository(db)
	articlesService := service.NewArticleService(authorsRepository)
	articlesHandler := handler.NewArticlesHandler(articlesService)

	v1.GET("/ping", func(c *gin.Context) {
		if err := db.Ping(); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"connection": err.Error(),
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"connection": "alive",
		})
	})

	articles := v1.Group("/articles")
	{
		articles.GET("/:id", articlesHandler.GetAuthorByID)
	}
}
