package routes

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/IloveNooodles/kumparan-techincal-test/internal/handler"
	"github.com/IloveNooodles/kumparan-techincal-test/internal/repository"
	"github.com/IloveNooodles/kumparan-techincal-test/internal/service"
	"github.com/gin-gonic/gin"
)

func wildcardRouting(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"connection": "alive",
		"time":       time.Now().Format(time.RFC1123),
	})
}

func RoutesInit(server *gin.Engine, db *sql.DB) {

	server.NoRoute(wildcardRouting)
	api := server.Group("/api")
	v1 := api.Group("/v1")

	authorsRepository := repository.NewAuthorsRepository(db)
	articlesRepository := repository.NewArticlesRepository(db)
	articlesService := service.NewArticleService(authorsRepository, articlesRepository)
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
		articles.GET("/", articlesHandler.GetArticles)
	}
}
