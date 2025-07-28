package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/michaeltukdev/Potok/internal/config"
	"github.com/spf13/cobra"
	"github.com/zalando/go-keyring"
)

var uploadFileCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload a file into your vault",
	Run: func(cmd *cobra.Command, args []string) {
		// Store locally on API key handling or remove {user} altogether
		user := "Michael"

		// Yes.
		vault := "Test"

		// Yes since remote path would include ther est
		remotePath := "Test.jpg"

		cfg, err := config.Load()
		if err != nil {
			fmt.Println("Error loading config:", err)
			return
		}

		url := fmt.Sprintf(cfg.APIURL+"/users/%s/vaults/%s/files/%s", user, vault, remotePath)
		localFile := "/home/athena/me/Test2.jpg"

		file, err := os.Open(localFile)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		req, err := http.NewRequest("POST", url, file)
		if err != nil {
			panic(err)
		}

		fmt.Println(url)

		secret, err := keyring.Get("potok", "api-key")
		if err != nil {
			fmt.Println("Error retrieving API key:", err)
			return
		}

		req.Header.Set("Authorization", secret)

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		fmt.Println("Status:", resp.Status)
		if resp.StatusCode != http.StatusCreated {
			fmt.Println("Upload failed!")
		}

		fmt.Println("Upload succeeded!")
	},
}
