package storage

import (
	"bytes"
	"fmt"
	"net/http"
)

func UploadFile(apiURL, username, vaultName, relPath string, encrypted []byte, apiKey string) error {
	url := fmt.Sprintf("%s/users/%s/vaults/%s/files/%s", apiURL, username, vaultName, relPath)
	req, err := http.NewRequest("POST", url, bytes.NewReader(encrypted))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", apiKey)
	req.Header.Set("Content-Type", "application/octet-stream")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("upload failed for %s: %s", relPath, resp.Status)
	}
	return nil
}
