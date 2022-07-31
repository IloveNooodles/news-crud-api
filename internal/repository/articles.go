package repository

import (
	"database/sql"

	"github.com/IloveNooodles/kumparan-techincal-test/internal/schema"
	"github.com/rs/zerolog/log"
)

type IArticlesRepository interface {
	GetArticles(query string, author string) ([]schema.Articles, error)
}
type articlesRepository struct {
	db *sql.DB
}

func (r *articlesRepository) GetArticles(query string, author string) ([]schema.Articles, error) {
	var listArticle []schema.Articles
	statement := `SELECT * FROM articles WHERE`

	if author != "" {
		statement += ` author_id = $1`
	}

	if query != "" {
		statement += ` lower(title) like lower('%$2%') or lower(body) like lower('%$2%')`
	}

	rows, err := r.db.Query(statement, author, query)
	log.Info().Msg(statement)

	if err != nil {
		log.Error().Msg("Error on querying")
		return listArticle, err
	}

	defer rows.Close()

	for rows.Next() {
		article := schema.Articles{}
		err := rows.Scan(&article.ID, &article.Author_ID, &article.Title, &article.Body, &article.Created_at)

		if err != nil {
			log.Error().Msg("Error on fetching data")
			rows.Close()
			return listArticle, err
		}

		listArticle = append(listArticle, article)

	}

	if err = rows.Err(); err != nil {
		log.Error().Msg("Error on getting rows")
		return listArticle, err
	}

	return listArticle, err
}

func NewArticlesRepository(db *sql.DB) IArticlesRepository {
	return &articlesRepository{
		db: db,
	}
}
