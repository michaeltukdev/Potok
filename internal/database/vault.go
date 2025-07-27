package database

import (
	"errors"
	"time"
)

type Vault struct {
	ID        int       `db:"id" json:"id"`
	UserID    int       `db:"user_id" json:"user_id"`
	Name      string    `db:"name" json:"name"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

func FetchUserVaults(apiKey string) ([]Vault, error) {
	user, err := FindByAPIKey(apiKey)
	if err != nil {
		return nil, errors.New("user not found")
	}

	rows, err := DB.Query("SELECT id, user_id, name, created_at, updated_at FROM vaults WHERE user_id = ?", user.id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var vaults []Vault
	for rows.Next() {
		var v Vault
		if err := rows.Scan(&v.ID, &v.UserID, &v.Name, &v.CreatedAt, &v.UpdatedAt); err != nil {
			return nil, err
		}
		vaults = append(vaults, v)
	}

	if vaults == nil {
		vaults = []Vault{}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return vaults, nil
}
