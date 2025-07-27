package database

import (
	"database/sql"
	"errors"
)

type User struct {
	id       int
	username string
	api_key  string
}

func FindByAPIKey(apiKey string) (*User, error) {
	row := DB.QueryRow("SELECT id, username, api_key FROM users WHERE api_key = ?", apiKey)

	var user User
	err := row.Scan(&user.id, &user.username, &user.api_key)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}
