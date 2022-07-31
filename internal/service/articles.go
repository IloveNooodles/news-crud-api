package service

import (
	"github.com/IloveNooodles/kumparan-techincal-test/internal/repository"
	"github.com/IloveNooodles/kumparan-techincal-test/internal/schema"
)

type IArticleService interface {
	GetAuthorByID(ID string) (schema.Author, error)
}

type articleService struct {
	authorsRepository repository.IAuthorsRepository
}

func (s *articleService) GetAuthorByID(ID string) (schema.Author, error) {
	author, err := s.authorsRepository.GetAuthorByID(ID)
	return author, err
}

func NewArticleService(authorsRepository repository.IAuthorsRepository) IArticleService {
	return &articleService{
		authorsRepository: authorsRepository,
	}
}
