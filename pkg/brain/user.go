package brain

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
)

type User struct {
	Username string `db:"username"`
}

func (b *Brain) CreateUser(username string) (*User, error) {
	user := &User{Username: username}
	_, err := b.db.NamedExec(`INSERT INTO users (username) VALUES (:username)`, user)
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			return nil, fmt.Errorf("user already exists: %w", ErrAlreadyExists)
		}
		return nil, err
	}
	return user, nil
}

func (b *Brain) GetUser(username string) (*User, error) {
	var user User
	if err := b.db.Get(&user, `SELECT * FROM users WHERE username = ?`, username); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("user does not exist: %w", ErrNotFound)
		}
		return nil, err
	}
	return &user, nil
}
