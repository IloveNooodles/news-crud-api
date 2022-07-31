package lib

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

var DB *sql.DB

func Init(url string) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		panic(err)
	}
	log.Info().Msg("Successfully connected to the database")
	DB = db
}
