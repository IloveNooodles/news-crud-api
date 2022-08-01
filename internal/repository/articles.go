package repository

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/IloveNooodles/kumparan-techincal-test/internal/schema"
	"github.com/rs/zerolog/log"
)

type IArticlesRepository interface {
	GetArticles(query string, author string, page int) ([]schema.ArticlesAuthor, error)
	CreateNewArticle(schema schema.Articles) error
}
type articlesRepository struct {
	db *sql.DB
}

func (r *articlesRepository) GetArticles(query string, author string, page int) ([]schema.ArticlesAuthor, error) {
	var listArticle []schema.ArticlesAuthor
	var rows *sql.Rows
	var err error
	statement := `SELECT * FROM articles NATURAL JOIN authors`
	lowerAuthor := strings.ToLower(author)
	lowerQuery := fmt.Sprintf(`%%%v%%`, strings.ToLower(query))
	LIMIT := 20
	offset := (page - 1) * LIMIT

	if author != "" && query != "" {
		statement += ` WHERE lower(name) = $1 AND (lower(title) like $2 OR lower(body) like $2) LIMIT 20 OFFSET $3 order by created_at desc`
		rows, err = r.db.Query(statement, lowerAuthor, lowerQuery, offset)
	} else if author != "" {
		statement += ` WHERE lower(name) = $1 LIMIT 20 OFFSET $2 order by created_at desc`
		rows, err = r.db.Query(statement, lowerAuthor, offset)
	} else if query != "" {
		statement += ` WHERE lower(title) like $1 OR lower(body) like $1 LIMIT 20 OFFSET $2 order by created_at desc`
		rows, err = r.db.Query(statement, lowerQuery, offset)
	} else {
		statement += ` LIMIT 20 OFFSET $1 order by created_at desc`
		rows, err = r.db.Query(statement, offset)
	}

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

func (r *articlesRepository) CreateNewArticle(schema schema.Articles) error {
	stmt, err := r.db.Prepare("insert into articles (id, author_id, title, body) values ($1, $2, $3, $4)")
	if err != nil {
		log.Error().Msg("Error on preparing sql statement")
		return err
	}

	_, err = stmt.Exec(schema.ID, schema.Author_ID, schema.Title, schema.Body)
	if err != nil {
		log.Error().Msg("Error on inserting data")
		return err
	}

	return nil
}

func NewArticlesRepository(db *sql.DB) IArticlesRepository {
	return &articlesRepository{
		db: db,
	}
}
