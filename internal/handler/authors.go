package handler

import (
	"net/http"

	"github.com/IloveNooodles/kumparan-techincal-test/internal/service"
	"github.com/gin-gonic/gin"
)

type IAuthorHandler interface {
	GetAuthors(c *gin.Context)
	CreateNewAuthor(c *gin.Context)
}

type authorHandler struct {
	authorService service.IAuthorService
}

func (h *authorHandler) GetAuthors(c *gin.Context) {
	authors, err := h.authorService.GetAuthors()
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"success": false,
			"message": "error when fetching data",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    authors,
	})

}

func (h *authorHandler) CreateNewAuthor(c *gin.Context) {

}

func NewAuthorHandler(authorService service.IAuthorService) IAuthorHandler {
	return &authorHandler{
		authorService: authorService,
	}
}
