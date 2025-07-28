package main

import (
	"fmt"

	"github.com/michaeltukdev/Potok/internal/config"
	"github.com/michaeltukdev/Potok/internal/prompt"
	"github.com/spf13/cobra"
)

func checkVault() {
	fmt.Println("False")
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

		fmt.Println(cfg)

		fmt.Println("Adding Vault...")

		vaultPath := prompt.Input("Path to your vault: ")
		vaultName := prompt.Input("Name your vault: ")
		checkVault()

		vaultPassword := prompt.Input("Encryption password: ")

		fmt.Printf("Vault Path: %s\nVault Name: %s\n", vaultPath, vaultName)
		fmt.Println(vaultPassword)
	},
}
