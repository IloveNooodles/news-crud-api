package service

import (
	"github.com/IloveNooodles/kumparan-techincal-test/internal/repository"
	"github.com/IloveNooodles/kumparan-techincal-test/internal/schema"
)

type IAuthorService interface {
	GetAuthors() ([]schema.Author, error)
	CreateNewAuthor(id string, name string) error
}

type authorService struct {
	authorsRepository repository.IAuthorsRepository
}

func (s *authorService) GetAuthors() ([]schema.Author, error) {
	authors, err := s.authorsRepository.GetAuthors()
	return authors, err
}

func (s *authorService) CreateNewAuthor(id string, name string) error {
	err := s.authorsRepository.CreateNewAuthor(id, name)
	return err
}

func NewAuthorService(authorRepository repository.IAuthorsRepository) IAuthorService {
	return &authorService{
		authorsRepository: authorRepository,
	}
}
