package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/michaeltukdev/Potok/internal/config"
	"github.com/spf13/cobra"
)

var setApiUrlCmd = &cobra.Command{
	Use:   "set-api-url",
	Short: "Set the Potok server API URL",
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter your Potok server URL (e.g., http://localhost:8080): ")

		url, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Failed to read input:", err)
			return
		}

		url = strings.TrimSpace(url)
		if url == "" {
			fmt.Println("API URL cannot be empty.")
			return
		}

		cfg, err := config.Load()
		if err != nil {
			fmt.Println("Failed to load config:", err)
			return
		}

		cfg.APIURL = url
		if err := config.Save(cfg); err != nil {
			fmt.Println("Failed to save config:", err)
			return
		}

		fmt.Println("Successfully saved your API URL!")
	},
}
