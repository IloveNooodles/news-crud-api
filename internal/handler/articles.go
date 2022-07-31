package handler

import (
	"net/http"

	"github.com/IloveNooodles/kumparan-techincal-test/internal/schema"
	"github.com/IloveNooodles/kumparan-techincal-test/internal/service"
	"github.com/gin-gonic/gin"
)

type IArticlesHandler interface {
	CreateNewArticle(c *gin.Context)
	GetArticles(c *gin.Context)
}

type articlesHandler struct {
	articleService service.IArticleService
}

func (h *articlesHandler) CreateNewArticle(c *gin.Context) {
	var json schema.Articles
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "error when parsing data, make sure to fill the data correctly",
		})
		return
	}

	err := h.articleService.CreateNewArticle(json)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"success": false,
			"message": "id or author_id is already exists",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "successfully created new articles",
	})
}

func (h *articlesHandler) GetArticles(c *gin.Context) {
	query, _ := c.GetQuery("query")
	author, _ := c.GetQuery("author")

	listOfArticles, err := h.articleService.GetArticles(query, author)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "error when fetching data",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    listOfArticles,
	})
}

func NewArticlesHandler(articleService service.IArticleService) IArticlesHandler {
	return &articlesHandler{
		articleService: articleService,
	}
}
