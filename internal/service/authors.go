package service

import (
	"github.com/ilovenooodles/news-crud-api/internal/repository"
	"github.com/ilovenooodles/news-crud-api/internal/schema"
)

type IAuthorService interface {
	GetAuthors(page int) ([]schema.Author, error)
	CreateNewAuthor(schema schema.Author) error
	UpdateAuthor(schema schema.Author) error
	DeleteAuthor(id string) error
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

func (s *authorService) UpdateAuthor(schema schema.Author) error {
	err := s.authorsRepository.UpdateAuthor(schema)
	return err
}

func (s *authorService) DeleteAuthor(id string) error {
	err := s.authorsRepository.DeleteAuthor(id)
	return err
}

func NewAuthorService(authorRepository repository.IAuthorsRepository) IAuthorService {
	return &authorService{
		authorsRepository: authorRepository,
	}
}
