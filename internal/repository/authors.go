package repository

import (
	"database/sql"

	"github.com/IloveNooodles/kumparan-techincal-test/internal/schema"
	"github.com/rs/zerolog/log"
)

type IAuthorsRepository interface {
	GetAuthorByID(ID string) (schema.Author, error)
}

type authorsRepository struct {
	db *sql.DB
}

func (r *authorsRepository) GetAuthorByID(ID string) (schema.Author, error) {
	stmt, err := r.db.Prepare("SELECT * FROM authors WHERE id = ?")
	if err != nil {
		log.Error().Msg("Error preparing sql statement")
		return schema.Author{}, err
	}
	res := schema.Author{}

	err = stmt.QueryRow(ID).Scan(res.ID, res.Name)

	if err != nil {
		log.Error().Msg("No users found")
		return schema.Author{}, err
	}

	return res, nil
}

func NewAuthorsRepository(db *sql.DB) IAuthorsRepository {
	return &authorsRepository{
		db: db,
	}
}
