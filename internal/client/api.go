package client

import (
	"encoding/json"
	"net/http"
	"time"
)

type Vault struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func MakeAuthenticatedRequest(apiKey, url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", apiKey)
	client := &http.Client{}
	return client.Do(req)
}

func MakeTestAuthenticatedRequest(apiKey, url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", apiKey)
	client := &http.Client{}
	return client.Do(req)
}

func ReadVaultsFromResponse(resp *http.Response) ([]Vault, error) {
	defer resp.Body.Close()
	var vaults []Vault
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&vaults); err != nil {
		return nil, err
	}
	return vaults, nil
}
