package main

import (
	"github.com/spf13/cobra"
)

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sync your vaults",
	Run: func(cmd *cobra.Command, args []string) {
		// Does something
	},
}
