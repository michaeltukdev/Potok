package client

import (
	"fmt"

	"github.com/zalando/go-keyring"
)

func VaultNameExists(vaults []Vault, name string) bool {
	for _, v := range vaults {
		if v.Name == name {
			return true
		}
	}
	return false
}

func CheckVault(api, username, vaultName string) (bool, error) {
	secret, err := keyring.Get("potok", "api-key")
	if err != nil {
		return false, fmt.Errorf("error retrieving API key: %w", err)
	}

	url := fmt.Sprintf("%s/users/%s/vaults", api, username)
	resp, err := MakeAuthenticatedRequest(secret, url)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return false, fmt.Errorf("unauthenticated! Please set or update your API key")
	}

	vaults, err := ReadVaultsFromResponse(resp)
	if err != nil {
		return false, err
	}

	return VaultNameExists(vaults, vaultName), nil
}
