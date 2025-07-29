package main

import (
	"fmt"
	"net/http"

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

		secret, err := keyring.Get("potok", "api-key")
		if err != nil {
			fmt.Println("Error retrieving API key:", err)
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

		fmt.Printf("Vault Path: %s\nVault Name: %s\n", vaultPath, vaultName)

		url := fmt.Sprintf("%s/users/%s/vaults/%s", cfg.APIURL, cfg.Username, vaultName)
		req, err := http.NewRequest("POST", url, nil)
		if err != nil {
			fmt.Println("Failed to create request:", err)
			return
		}
		req.Header.Set("Authorization", secret)

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Println("Failed to register vault:", err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusConflict {
			fmt.Printf("A vault named '%s' already exists. Please choose a different name.\n", vaultName)
			return
		}
		if resp.StatusCode != http.StatusCreated {
			fmt.Printf("Failed to register vault! Server returned: %s\n", resp.Status)
			return
		}

		fmt.Println("Vault registered successfully!")

		vaultInfo := config.VaultInfo{
			Name: vaultName,
			Path: vaultPath,
		}

		cfg.AddVault(vaultInfo)
		if err := config.Save(cfg); err != nil {
			fmt.Println("Failed to save vault info locally:", err)
			return
		}

		fmt.Println("Vault info saved locally!")
	},
}
