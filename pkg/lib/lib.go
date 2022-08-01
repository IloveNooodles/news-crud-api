package lib

import (
	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

var DB *sql.DB

func Init(url string) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		panic(err)
	}

  driver, err := postgres.WithInstance(db, &postgres.Config{})
  m, err := migrate.NewWithDatabaseInstance(
    
  )

	log.Info().Msg("Successfully connected to the database")
	DB = db
}
