package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"io/ioutil"
)

// DecryptFile decrypts the input file and writes the result to output file.
func DecryptFile(password, inputPath, outputPath string) error {
	data, err := ioutil.ReadFile(inputPath)
	if err != nil {
		return err
	}

	if len(data) < saltSize+nonceSize {
		return fmt.Errorf("file too short")
	}

	salt := data[:saltSize]
	nonce := data[saltSize : saltSize+nonceSize]
	ciphertext := data[saltSize+nonceSize:]

	key, err := DeriveKey([]byte(password), salt)
	if err != nil {
		return err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(outputPath, plaintext, 0644)
}
