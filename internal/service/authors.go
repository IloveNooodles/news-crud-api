package service

import (
	"github.com/IloveNooodles/kumparan-techincal-test/internal/repository"
	"github.com/IloveNooodles/kumparan-techincal-test/internal/schema"
)

type IAuthorService interface {
	GetAuthors(page int) ([]schema.Author, error)
	CreateNewAuthor(schema schema.Author) error
}

type authorService struct {
	authorsRepository repository.IAuthorsRepository
}

func (s *authorService) GetAuthors(page int) ([]schema.Author, error) {
	authors, err := s.authorsRepository.GetAuthors(page)
	return authors, err
}

func (s *authorService) CreateNewAuthor(schema schema.Author) error {
	err := s.authorsRepository.CreateNewAuthor(schema)
	return err
}

func NewAuthorService(authorRepository repository.IAuthorsRepository) IAuthorService {
	return &authorService{
		authorsRepository: authorRepository,
	}
}