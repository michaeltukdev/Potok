package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/zalando/go-keyring"
)

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sync your vaults",
	Run: func(cmd *cobra.Command, args []string) {
		secret, err := keyring.Get("potok", "api-key")
		if err != nil {
			fmt.Println("Error retrieving API key:", err)
			return
		}

		fmt.Println("API key:", secret)
	},
}
