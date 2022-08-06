package repository

import (
	"database/sql"

	"github.com/ilovenooodles/news-crud-api/internal/schema"
	"github.com/rs/zerolog/log"
)

type IAuthorsRepository interface {
	GetAuthors(page int) ([]schema.Author, error)
	CreateNewAuthor(schema schema.Author) error
	UpdateAuthor(schema schema.Author) error
	DeleteAuthor(id string) error
}

type authorsRepository struct {
	db *sql.DB
}

func (r *authorsRepository) GetAuthors(page int) ([]schema.Author, error) {
	var listAuthor []schema.Author
	LIMIT := 20
	offset := (page - 1) * LIMIT
	rows, err := r.db.Query("select * from authors LIMIT 20 offset $1", offset)
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

func (r *authorsRepository) CreateNewAuthor(schema schema.Author) error {
	stmt, err := r.db.Prepare("insert into authors (id, name) values ($1, $2)")
	if err != nil {
		log.Error().Msg("Error preparing sql statement")
		return err
	}

	_, err = stmt.Exec(schema.ID, schema.Name)
	if err != nil {
		log.Error().Msg("Error when inserting to database")
		return err
	}

	return nil
}

func (r *authorsRepository) UpdateAuthor(schema schema.Author) error {
	stmt, err := r.db.Prepare("update authors set name = $1 where id = $2")
	if err != nil {
		log.Error().Msg("Error preparing sql statement")
		return err
	}

	_, err = stmt.Exec(schema.Name, schema.ID)

	if err != nil {
		log.Error().Msg("Authors have post that cannot be deleted")
		return err
	}

	return nil
}

func (r *authorsRepository) DeleteAuthor(id string) error {
	stmt, err := r.db.Prepare("delete from authors where id = $1")

	if err != nil {
		log.Error().Msg("Error preparing sql statement")
		return err
	}

	_, err = stmt.Exec(id)

	if err != nil {
		log.Error().Msg("Authors have post that cannot be deleted")
		return err
	}

	return nil
}

func NewAuthorsRepository(db *sql.DB) IAuthorsRepository {
	return &authorsRepository{
		db: db,
	}
}
