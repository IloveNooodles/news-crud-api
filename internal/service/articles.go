package service

import (
	"github.com/IloveNooodles/kumparan-techincal-test/internal/repository"
	"github.com/IloveNooodles/kumparan-techincal-test/internal/schema"
)

type IArticleService interface {
	CreateNewArticle(schema schema.Articles) error
	GetArticles(query string, author string) ([]schema.ArticlesAuthor, error)
}

type articleService struct {
	articlesRepository repository.IArticlesRepository
}

func (s *articleService) GetArticles(query string, author string) ([]schema.ArticlesAuthor, error) {
	listOfAuthor, err := s.articlesRepository.GetArticles(query, author)
	return listOfAuthor, err
}

func (s *articleService) CreateNewArticle(schema schema.Articles) error {
	err := s.articlesRepository.CreateNewArticle(schema)
	return err
}

func NewArticleService(articlesRepository repository.IArticlesRepository) IArticleService {
	return &articleService{
		articlesRepository: articlesRepository,
	}
}
