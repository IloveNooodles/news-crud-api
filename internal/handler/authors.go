package handler

import (
	"net/http"
	"strconv"

	"github.com/IloveNooodles/kumparan-techincal-test/internal/schema"
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
	page, _ := c.GetQuery("page")
	var pageInt int

	if num, err := strconv.Atoi(page); err != nil {
		pageInt = 1
	} else {
		pageInt = num
	}

	authors, err := h.authorService.GetAuthors(pageInt)
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
	var json schema.Author
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "error when parsing data, fill the data correctly",
		})
		return
	}

	err := h.authorService.CreateNewAuthor(json)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"success": false,
			"message": "id is already exists",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "successfuly created new user",
	})
}

func NewAuthorHandler(authorService service.IAuthorService) IAuthorHandler {
	return &authorHandler{
		authorService: authorService,
	}
}
