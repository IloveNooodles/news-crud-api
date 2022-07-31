package service

import (
	"github.com/IloveNooodles/kumparan-techincal-test/internal/repository"
	"github.com/IloveNooodles/kumparan-techincal-test/internal/schema"
)

type IArticleService interface {
	GetAuthorByID(ID string) (schema.Author, error)
	GetArticles(query string, author string) ([]schema.ArticlesAuthor, error)
}

type articleService struct {
	authorsRepository  repository.IAuthorsRepository
	articlesRepository repository.IArticlesRepository
}

func (s *articleService) GetAuthorByID(ID string) (schema.Author, error) {
	author, err := s.authorsRepository.GetAuthorByID(ID)
	return author, err
}

func (s *articleService) GetArticles(query string, author string) ([]schema.ArticlesAuthor, error) {
	listOfAuthor, err := s.articlesRepository.GetArticles(query, author)
	return listOfAuthor, err
}

func NewArticleService(authorsRepository repository.IAuthorsRepository, articlesRepository repository.IArticlesRepository) IArticleService {
	return &articleService{
		authorsRepository:  authorsRepository,
		articlesRepository: articlesRepository,
	}
}
