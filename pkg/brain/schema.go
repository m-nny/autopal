package brain

import "github.com/jmoiron/sqlx"

const _schema = `
	CREATE TABLE IF NOT EXISTS users (
		username TEXT PRIMARY KEY,
	  balance INTEGER NOT NULL
	);
`

func initSchema(db *sqlx.DB) error {
	if _, err := db.Exec(_schema); err != nil {
		return err
	}
	return nil
}
