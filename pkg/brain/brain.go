package brain

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/tursodatabase/go-libsql"
)

type Brain struct {
	db *sqlx.DB
}

func New(dsn string) (*Brain, error) {
	db, err := sqlx.Open("libsql", dsn)
	if err != nil {
		return nil, err
	}
	if err := initSchema(db); err != nil {
		return nil, err
	}
	return &Brain{db: db}, nil
}

func (brain *Brain) Close() error {
	return brain.db.Close()
}
