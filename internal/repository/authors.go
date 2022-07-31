package repository

import (
	"database/sql"

	"github.com/IloveNooodles/kumparan-techincal-test/internal/schema"
	"github.com/rs/zerolog/log"
)

type IAuthorsRepository interface {
	GetAuthors() ([]schema.Author, error)
	CreateNewAuthor(id string, name string) error
}

type authorsRepository struct {
	db *sql.DB
}

func (r *authorsRepository) GetAuthors() ([]schema.Author, error) {
	var listAuthor []schema.Author
	rows, err := r.db.Query("select * from authors")
	if err != nil {
		log.Error().Msg("Error preparing sql statement")
		return listAuthor, err
	}

	defer rows.Close()

	for rows.Next() {
		author := schema.Author{}
		err := rows.Scan(&author.ID, &author.Name)

		if err != nil {
			log.Error().Msg("Error when getting data from sql")
			return listAuthor, err
		}
		listAuthor = append(listAuthor, author)
	}

	return listAuthor, nil
}

func (r *authorsRepository) CreateNewAuthor(id string, name string) error {
	stmt, err := r.db.Prepare("insert into authors (id, name) values ($1, $2)")
	if err != nil {
		log.Error().Msg("Error preparing sql statement")
		return err
	}

	_, err = stmt.Exec(id, name)
	if err != nil {
		log.Error().Msg("Error when inserting to database")
		return err
	}

	return nil
}

func NewAuthorsRepository(db *sql.DB) IAuthorsRepository {
	return &authorsRepository{
		db: db,
	}
}
