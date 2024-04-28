package brain

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
)

const (
	UserStartingBalance = 10
)

type User struct {
	Username string `db:"username"`
	Balance  int    `db:"balance"`
}

// newUser returns User with some default values
func newUser(username string) *User {
	return &User{
		Username: username,
		Balance:  UserStartingBalance,
	}
}

func (b *Brain) CreateUser(username string) (*User, error) {
	user := newUser(username)
	_, err := b.db.NamedExec(`INSERT INTO users (username, balance) VALUES (:username, :balance)`, user)
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
