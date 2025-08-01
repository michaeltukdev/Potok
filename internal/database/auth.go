package database

import (
	"database/sql"
	"errors"
)

type User struct {
	Id       int
	Username string
	Api_key  string
}

func FindByAPIKey(apiKey string) (*User, error) {
	row := DB.QueryRow("SELECT id, username, api_key FROM users WHERE api_key = ?", apiKey)

	var user User
	err := row.Scan(&user.Id, &user.Username, &user.Api_key)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}
