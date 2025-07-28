package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"os"

	"golang.org/x/crypto/scrypt"
)

const (
	saltSize  = 16
	nonceSize = 12
	keyLen    = 32 // AES-256
)

// Takes the 'password' from the user (currently hardcoded)
// Takes the generated 'salt' based on saltSize random integers
// The numbers I don't really understand
// Maximum length of key being a 32 byte string (256 bits)
func DeriveKey(password, salt []byte) ([]byte, error) {
	return scrypt.Key(password, salt, 1<<15, 8, 1, keyLen)
}

// EncryptFile encrypts the input file and writes the result to output file.
func EncryptFile(password, inputPath string) ([]byte, error) {

	// Read entire file and store it in 'plainText'
	f, err := os.Open(inputPath)
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	// Initialises variable 'salt' to `saltSize` of 0's
	salt := make([]byte, saltSize)
	// Uses `rand.Read` to initialise each 0 to a random number
	// Which is completely random each time
	if _, err := rand.Read(salt); err != nil {
		return nil, err
	}

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

	nonce := make([]byte, nonceSize)
	if _, err := rand.Read(nonce); err != nil {
		return nil, err
	}

	ciphertext := gcm.Seal(nil, nonce, data, nil)

	// Output: salt | nonce | ciphertext
	out := append(salt, nonce...)
	out = append(out, ciphertext...)

	return out, nil
}

// func main() {
// 	// Grab vault name (potok will always be the service)
// 	service := "potok"
// 	user := "my-vaults"

// 	// Take in password from user
// 	var password string
// 	fmt.Scan(&password)

// 	// Store password in the keyring service
// 	err := keyring.Set(service, user, password)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	secret, err := keyring.Get(service, user)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Call the EncryptFile
// 	err = EncryptFile(secret, "plain.txt", "encrypted.txt")
// 	if err != nil {
// 		fmt.Println("Encryption error:", err)
// 		return
// 	}
// 	fmt.Println("File encrypted.")

// 	// Instantly DecryptFile
// 	err = DecryptFile(secret, "encrypted.txt", "decrypted.txt")
// 	if err != nil {
// 		fmt.Println("Decryption error:", err)
// 		return
// 	}
// 	fmt.Println("File decrypted.")
// }
