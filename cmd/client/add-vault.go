package main

import (
	"fmt"

	"github.com/michaeltukdev/Potok/internal/client"
	"github.com/michaeltukdev/Potok/internal/config"
	"github.com/michaeltukdev/Potok/internal/prompt"
	"github.com/spf13/cobra"
	"github.com/zalando/go-keyring"
)

// TODO: Probably move from this file
func VaultNameExists(vaults []client.Vault, name string) bool {
	for _, v := range vaults {
		if v.Name == name {
			return true
		}
	}
	return false
}

// TODO: Probably move from this file
func checkVault(api, username, vaultName string) (bool, error) {
	secret, err := keyring.Get("potok", "api-key")
	if err != nil {
		return false, fmt.Errorf("error retrieving API key: %w", err)
	}

	url := fmt.Sprintf("%s/users/%s/vaults", api, username)
	resp, err := client.MakeAuthenticatedRequest(secret, url)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return false, fmt.Errorf("unauthenticated! Please set or update your API key")
	}

	vaults, err := client.ReadVaultsFromResponse(resp)
	if err != nil {
		return false, err
	}

	return VaultNameExists(vaults, vaultName), nil
}

var addVaultCmd = &cobra.Command{
	Use:   "add-vault",
	Short: "Select a Vault to be backed up securely!",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.MustLoadWithAPIURL()
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Adding Vault...")

		vaultPath := prompt.Input("Path to your vault: ")

		var vaultName string
		for {
			vaultName = prompt.Input("Name your vault: ")
			exists, err := checkVault(cfg.APIURL, cfg.Username, vaultName)
			if err != nil {
				fmt.Println("Error checking vault:", err)
				return
			}
			if exists {
				fmt.Printf("A vault named '%s' already exists. Please choose a different name.\n", vaultName)
			} else {
				break
			}
		}

		// vaultPassword := prompt.Input("Encryption password: ")

		fmt.Printf("Vault Path: %s\nVault Name: %s\n", vaultPath, vaultName)
	},
}
