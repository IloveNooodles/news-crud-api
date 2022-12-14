package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ilovenooodles/news-crud-api/internal/schema"
	"github.com/ilovenooodles/news-crud-api/internal/service"
	"github.com/ilovenooodles/news-crud-api/pkg/redis"
	"github.com/rs/zerolog/log"
)

type IAuthorHandler interface {
	GetAuthors(c *gin.Context)
	CreateNewAuthor(c *gin.Context)
	UpdateAuthor(c *gin.Context)
	DeleteAuthor(c *gin.Context)
}

type authorHandler struct {
	authorService service.IAuthorService
}

func (h *authorHandler) GetAuthors(c *gin.Context) {
	page, _ := c.GetQuery("page")
	var pageInt int
	var authors []schema.Author
	var err error

	if num, err := strconv.Atoi(page); err != nil {
		pageInt = 1
	} else {
		pageInt = num
	}

	rdb := redis.NewRedisClient()
	cachedData, err := rdb.Get(context.Background(), "authors").Bytes()

	if err != nil {
		authors, err = h.authorService.GetAuthors(pageInt)

		if err != nil {
			log.Info().Msg("failed to get from database")

			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "error when fetching data",
			})
			return
		}

		cachedData, err := json.Marshal(authors)

		if err != nil {
			log.Info().Msg("failed to convert to redis")

			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "error when fetching data",
			})
			return
		}

		err = rdb.Set(context.Background(), "authors", cachedData, time.Second*30).Err()

		if err != nil {
			log.Info().Msg("failed to SET to redis")
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "error when fetching data",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    authors,
		})
		return
	}

	err = json.Unmarshal(cachedData, &authors)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "error when fetching data",
		})
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

func (h *authorHandler) DeleteAuthor(c *gin.Context) {
	var json schema.AuthorRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"success": false,
			"message": "missing parameter",
		})
		return
	}

	err := h.authorService.DeleteAuthor(json.ID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "author is not found or have post",
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"success": true,
		"message": "successfully deleted author",
	})

}

func (h *authorHandler) UpdateAuthor(c *gin.Context) {
	var json schema.Author
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"success": false,
			"message": "missing parameter",
		})
		return
	}

	err := h.authorService.UpdateAuthor(json)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "no author is found",
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"success": true,
		"message": "author updated successfuly",
	})
}

func NewAuthorHandler(authorService service.IAuthorService) IAuthorHandler {
	return &authorHandler{
		authorService: authorService,
	}
}
