package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ilovenooodles/news-crud-api/internal/schema"
	"github.com/ilovenooodles/news-crud-api/internal/service"
	"github.com/ilovenooodles/news-crud-api/pkg/redis"
	"github.com/rs/zerolog/log"
)

type IArticlesHandler interface {
	GetArticles(c *gin.Context)
	CreateNewArticle(c *gin.Context)
	UpdateArticle(c *gin.Context)
	DeleteArticle(c *gin.Context)
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
	page, _ := c.GetQuery("page")
	var pageInt int
	var err error
	var listOfArticles []schema.ArticlesAuthor

	if num, err := strconv.Atoi(page); err != nil {
		pageInt = 1
	} else {
		pageInt = num
	}

	key := fmt.Sprintf("articles-q%v-a%v-p%v", query, author, page)

	rdb := redis.NewRedisClient()
	cachedData, err := rdb.Get(context.Background(), key).Bytes()
	if err != nil {
		listOfArticles, err = h.articleService.GetArticles(query, author, pageInt)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"message": "error when fetching data",
			})
			return
		}

		cahcedData, err := json.Marshal(listOfArticles)

		if err != nil {
			log.Info().Msg("failed to convert to redis")

			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "error when fetching data",
			})
			return
		}

		err = rdb.Set(context.Background(), key, cahcedData, time.Second*30).Err()

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
			"data":    listOfArticles,
		})
		return
	}

	err = json.Unmarshal(cachedData, &listOfArticles)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
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

func (h *articlesHandler) UpdateArticle(c *gin.Context) {
	var json schema.ArticlesRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"success": false,
			"message": "missing parameter",
		})
		return
	}

	err := h.articleService.UpdateArticle(json)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "no such id exists",
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"success": true,
		"message": "successfully update article",
	})
}

func (h *articlesHandler) DeleteArticle(c *gin.Context) {
	var json schema.ArticlesDeleteRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"success": false,
			"message": "missing parameter",
		})
		return
	}

	err := h.articleService.DeleteArticle(json.ID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "no such id exists",
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"success": true,
		"message": "successfully deleted article",
	})
}

func NewArticlesHandler(articleService service.IArticleService) IArticlesHandler {
	return &articlesHandler{
		articleService: articleService,
	}
}
