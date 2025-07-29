package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

func DecryptBytes(password string, data []byte) ([]byte, error) {
	if len(data) < saltSize+nonceSize {
		return nil, fmt.Errorf("file too short")
	}

	salt := data[:saltSize]
	nonce := data[saltSize : saltSize+nonceSize]
	ciphertext := data[saltSize+nonceSize:]

	key, err := DeriveKey([]byte(password), salt)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}
