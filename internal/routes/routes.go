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

	articlesRepository := repository.NewArticlesRepository(db)
	articlesService := service.NewArticleService(articlesRepository)
	articlesHandler := handler.NewArticlesHandler(articlesService)
	authorsRepository := repository.NewAuthorsRepository(db)
	authorsService := service.NewAuthorService(authorsRepository)
	authorsHandler := handler.NewAuthorHandler(authorsService)

	articles := v1.Group("/articles")
	{
		articles.GET("/", articlesHandler.GetArticles)
		articles.POST("/", articlesHandler.CreateNewArticle)
	}

	author := v1.Group("/authors")
	{
		author.GET("/", authorsHandler.GetAuthors)
		author.POST("/", authorsHandler.CreateNewAuthor)
	}
}
