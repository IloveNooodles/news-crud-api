package repository

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/ilovenooodles/news-crud-api/internal/schema"
	"github.com/rs/zerolog/log"
)

type IArticlesRepository interface {
	GetArticles(query string, author string, page int) ([]schema.ArticlesAuthor, error)
	CreateNewArticle(schema schema.Articles) error
	UpdateArticle(schema schema.ArticlesRequest) error
	DeleteArticle(id string) error
}
type articlesRepository struct {
	db *sql.DB
}

func (r *articlesRepository) GetArticles(query string, author string, page int) ([]schema.ArticlesAuthor, error) {
	var listArticle []schema.ArticlesAuthor
	var rows *sql.Rows
	var err error
	statement := `SELECT a.*, au."name" FROM articles a inner JOIN authors au on a.author_id = au.id`
	lowerAuthor := strings.ToLower(author)
	lowerQuery := fmt.Sprintf(`%%%v%%`, strings.ToLower(query))
	LIMIT := 20
	offset := (page - 1) * LIMIT

	if author != "" && query != "" {
		statement += ` WHERE lower(name) = $1 AND (lower(title) like $2 OR lower(body) like $2) ORDER BY created_at DESC LIMIT 20 OFFSET $3`
		rows, err = r.db.Query(statement, lowerAuthor, lowerQuery, offset)
	} else if author != "" {
		statement += ` WHERE lower(name) = $1 ORDER BY created_at DESC LIMIT 20 OFFSET $2`
		rows, err = r.db.Query(statement, lowerAuthor, offset)
	} else if query != "" {
		statement += ` WHERE lower(title) like $1 OR lower(body) like $1 ORDER BY created_at DESC LIMIT 20 OFFSET $2`
		rows, err = r.db.Query(statement, lowerQuery, offset)
	} else {
		statement += ` ORDER BY created_at DESC LIMIT 20 OFFSET $1`
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

func (r *articlesRepository) UpdateArticle(schema schema.ArticlesRequest) error {
	stmt, err := r.db.Prepare("update articles set title = $1, body = $2, created_at = now() where id = $3")
	if err != nil {
		log.Error().Msg("Error on preparing sql statement")
		return err
	}

	_, err = stmt.Exec(schema.Title, schema.Body, schema.ID)
	if err != nil {
		log.Error().Msg("Error on updating data")
		return err
	}

	return nil
}

func (r *articlesRepository) DeleteArticle(id string) error {
	stmt, err := r.db.Prepare("delete from articles where id = $1")
	if err != nil {
		log.Error().Msg("Error on preparing sql statement")
		return err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		log.Error().Msg("Error on deleting data")
		return err
	}

	return nil
}

func NewArticlesRepository(db *sql.DB) IArticlesRepository {
	return &articlesRepository{
		db: db,
	}
}
