package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func promptInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		input, err := reader.ReadString('\n')

		if err != nil {
			log.Fatalf("Failed to read input: %v", err)
		}

		input = strings.TrimSpace(input)

		if input == "" {
			fmt.Println("Input cannot be empty.")
			continue
		}

		return input
	}
}

var addVaultCmd = &cobra.Command{
	Use:   "add-vault",
	Short: "Select a Vault to be backed up securely!",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Adding Vault...")

		vaultPath := promptInput("Path to your vault: ")
		vaultName := promptInput("Name your vault: ")
		vaultPassword := promptInput("Encryption password: ")

		fmt.Printf("Vault Path: %s\nVault Name: %s\n", vaultPath, vaultName)
		fmt.Println(vaultPassword)
	},
}
