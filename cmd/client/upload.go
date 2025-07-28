package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"

	"github.com/michaeltukdev/Potok/internal/config"
	"github.com/michaeltukdev/Potok/internal/crypto"
	"github.com/spf13/cobra"
	"github.com/zalando/go-keyring"
)

var uploadFileCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload a file into your vault",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.MustLoadWithAPIURL()
		if err != nil {
			fmt.Println(err)
			return
		}

		// Yes.
		vault := "Not"

		// Yes since remote path would include ther est
		remotePath := "Test.jpg"

		url := fmt.Sprintf(cfg.APIURL+"/users/%s/vaults/%s/files/%s", cfg.Username, vault, remotePath)
		localFile := "/home/athena/me/Test2.jpg"

		file, err := os.Open(localFile)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		// TODO: Update password
		encryptedBytes, err := crypto.EncryptFile("123", localFile)
		if err != nil {
			fmt.Println("Encryption error:", err)
			return
		}

		req, err := http.NewRequest("POST", url, bytes.NewReader(encryptedBytes))
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
			return
		}

		fmt.Println("Upload succeeded!")
	},
}
