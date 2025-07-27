package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/michaeltukdev/Potok/internal/client"
	"github.com/michaeltukdev/Potok/internal/config"
	"github.com/spf13/cobra"
	"github.com/zalando/go-keyring"
)

var setApiKeyCmd = &cobra.Command{
	Use:   "set-api-key",
	Short: "Set your API key",
	Run: func(cmd *cobra.Command, args []string) {
		service := "potok"
		user := "api-key"

		cfg, err := config.Load()
		if err != nil {
			fmt.Println("Error loading config:", err)
			return
		}

		if cfg.APIURL == "" {
			fmt.Println("API URL is not set. Please run 'potok set-api-url' first.")
			return
		}

		fmt.Print("Please enter your API key: ")

		var apiKey string
		_, err = fmt.Scanln(&apiKey)
		if err != nil {
			log.Fatal("Failed to read input:", err)
		}

		apiKey = strings.TrimSpace(apiKey)
		if apiKey == "" {
			log.Fatal("API key cannot be empty.")
		}

		r, err := client.MakeAuthenticatedRequest(apiKey, cfg.APIURL+"/users/1/vaults")
		if err != nil {
			log.Fatal(err)
		}

		if r.StatusCode != 200 {
			fmt.Println("Authentication Test Request Failed! Key not modified or saved. Please check your API url and make sure the server is running!")
			return
		}

		fmt.Println("Authentication Test Request Success!")

		err = keyring.Set(service, user, apiKey)
		if err != nil {
			log.Fatal("Failed to save API key:", err)
		}

		fmt.Println("Successfully saved your API key!")
	},
}
