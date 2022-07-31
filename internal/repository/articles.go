package repository

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/IloveNooodles/kumparan-techincal-test/internal/schema"
	"github.com/rs/zerolog/log"
)

type IArticlesRepository interface {
	GetArticles(query string, author string) ([]schema.ArticlesAuthor, error)
}
type articlesRepository struct {
	db *sql.DB
}

func (r *articlesRepository) GetArticles(query string, author string) ([]schema.ArticlesAuthor, error) {
	var listArticle []schema.ArticlesAuthor
	statement := `SELECT * FROM articles NATURAL JOIN authors`

	var rows *sql.Rows
	var err error
	lowerAuthor := strings.ToLower(author)
	lowerQuery := fmt.Sprintf(`%%%v%%`, strings.ToLower(query))
	fmt.Print(lowerQuery)

	if author != "" && query != "" {
		statement += ` WHERE lower(name) = $1 AND (lower(title) like $2 OR lower(body) like $2)`
		rows, err = r.db.Query(statement, lowerAuthor, lowerQuery)
	} else if author != "" {
		statement += ` WHERE lower(name) = $1`
		rows, err = r.db.Query(statement, lowerAuthor)
	} else if query != "" {
		statement += ` WHERE lower(title) like $1 OR lower(body) like $1`
		rows, err = r.db.Query(statement, lowerQuery)
	} else {
		rows, err = r.db.Query(statement)
	}

	log.Info().Msg(statement)

	if err != nil {
		log.Error().Msg("Error on querying")
		return listArticle, err
	}

	defer rows.Close()

	for rows.Next() {
		article := schema.ArticlesAuthor{}
		err := rows.Scan(&article.ID, &article.Author_ID, &article.Title, &article.Body, &article.Created_at, &article.Name)

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
