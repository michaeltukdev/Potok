package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
	"github.com/zalando/go-keyring"
)

var setApiKeyCmd = &cobra.Command{
	Use:   "set-api-key",
	Short: "Set your API key",
	Run: func(cmd *cobra.Command, args []string) {
		service := "potok"
		user := "api-key"

		fmt.Print("Please enter your API key: ")

		var apiKey string
		_, err := fmt.Scanln(&apiKey)
		if err != nil {
			log.Fatal("Failed to read input:", err)
		}

		apiKey = strings.TrimSpace(apiKey)
		if apiKey == "" {
			log.Fatal("API key cannot be empty.")
		}

		err = keyring.Set(service, user, apiKey)
		if err != nil {
			log.Fatal("Failed to save API key:", err)
		}

		fmt.Println("Successfully saved your API key!")
	},
}
