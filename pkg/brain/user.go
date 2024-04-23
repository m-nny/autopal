package brain

import "context"

type User struct {
	Username string `db:"username"`
}

func (b *Brain) UpsertUser(ctx context.Context, username string) (*User, error) {
	user := &User{Username: username}
	if _, err := b.db.NamedExecContext(ctx, `
			INSERT INTO users (username) VALUES (:username)
		  ON CONFLICT DO NOTHING
		`, user); err != nil {
		return nil, err
	}
	return user, nil
}
