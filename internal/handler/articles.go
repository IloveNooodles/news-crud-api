package handler

import (
	"net/http"

	"github.com/IloveNooodles/kumparan-techincal-test/internal/service"
	"github.com/gin-gonic/gin"
)

type IArticlesHandler interface {
	GetAuthorByID(c *gin.Context)
}

type articlesHandler struct {
	articleService service.IArticleService
}

func (h *articlesHandler) GetAuthorByID(c *gin.Context) {
	id := c.Param("id")
	author, err := h.articleService.GetAuthorByID(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"author":  author,
	})
}

func (h *articlesHandler) GetArticles(c *gin.Context) {
	query, _ := c.GetQuery("query")
	author, _ := c.GetQuery("author")

	listOfArticles, err := h.articleService.GetArticles(query, author)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":  true,
		"articles": listOfArticles,
	})
}

func NewArticlesHandler(articleService service.IArticleService) IArticlesHandler {
	return &articlesHandler{
		articleService: articleService,
	}
}
