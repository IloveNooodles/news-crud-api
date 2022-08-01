package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/IloveNooodles/kumparan-techincal-test/internal/schema"
	"github.com/IloveNooodles/kumparan-techincal-test/internal/service"
	"github.com/IloveNooodles/kumparan-techincal-test/pkg/redis"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
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
	page, _ := c.GetQuery("page")
	var pageInt int
	var err error
	var listOfArticles []schema.ArticlesAuthor

	if num, err := strconv.Atoi(page); err != nil {
		pageInt = 1
	} else {
		pageInt = num
	}

	rdb := redis.NewRedisClient()
	cachedData, err := rdb.Get(context.Background(), "articles").Bytes()

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

		err = rdb.Set(context.Background(), "articles", cahcedData, time.Hour).Err()

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

func NewArticlesHandler(articleService service.IArticleService) IArticlesHandler {
	return &articlesHandler{
		articleService: articleService,
	}
}
