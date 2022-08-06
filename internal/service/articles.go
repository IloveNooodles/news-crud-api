package service

import (
	"github.com/IloveNooodles/kumparan-techincal-test/internal/repository"
	"github.com/IloveNooodles/kumparan-techincal-test/internal/schema"
)

type IArticleService interface {
	GetArticles(query string, author string, page int) ([]schema.ArticlesAuthor, error)
	CreateNewArticle(schema schema.Articles) error
	UpdateArticle(schema schema.ArticlesRequest) error
	DeleteArticle(id string) error
}

type articleService struct {
	articlesRepository repository.IArticlesRepository
}

func (s *articleService) GetArticles(query string, author string, page int) ([]schema.ArticlesAuthor, error) {
	listOfAuthor, err := s.articlesRepository.GetArticles(query, author, page)
	return listOfAuthor, err
}

func (s *articleService) CreateNewArticle(schema schema.Articles) error {
	err := s.articlesRepository.CreateNewArticle(schema)
	return err
}

func (s *articleService) UpdateArticle(schema schema.ArticlesRequest) error {
	err := s.articlesRepository.UpdateArticle(schema)
	return err
}

func (s *articleService) DeleteArticle(id string) error {
	err := s.articlesRepository.DeleteArticle(id)
	return err
}

func NewArticleService(articlesRepository repository.IArticlesRepository) IArticleService {
	return &articleService{
		articlesRepository: articlesRepository,
	}
}
