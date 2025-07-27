package main

func main() {
	Execute()
}

// import (
// 	"crypto/aes"
// 	"crypto/cipher"
// 	"crypto/rand"
// 	"fmt"
// 	"io/ioutil"

// 	"golang.org/x/crypto/scrypt"
// )

// const (
// 	saltSize  = 16
// 	nonceSize = 12
// 	keyLen    = 32 // AES-256
// )

// // Takes the 'password' from the user (currently hardcoded)
// // Takes the generated 'salt' based on saltSize random integers
// // The numbers I don't really understand
// // Maximum length of key being a 32 byte string (256 bits)
// func DeriveKey(password, salt []byte) ([]byte, error) {
// 	return scrypt.Key(password, salt, 1<<15, 8, 1, keyLen)
// }

// // EncryptFile encrypts the input file and writes the result to output file.
// func EncryptFile(password, inputPath, outputPath string) error {
// 	// Read entire file and store it in 'plainText'
// 	plaintext, err := ioutil.ReadFile(inputPath)
// 	if err != nil {
// 		return err
// 	}

// 	// Initialises variable 'salt' to `saltSize` of 0's
// 	salt := make([]byte, saltSize)
// 	// Uses `rand.Read` to initialise each 0 to a random number
// 	// Which is completely random each time
// 	if _, err := rand.Read(salt); err != nil {
// 		return err
// 	}
// 	fmt.Println(salt)

// 	key, err := DeriveKey([]byte(password), salt)
// 	if err != nil {
// 		return err
// 	}

// 	block, err := aes.NewCipher(key)
// 	if err != nil {
// 		return err
// 	}

// 	gcm, err := cipher.NewGCM(block)
// 	if err != nil {
// 		return err
// 	}

// 	nonce := make([]byte, nonceSize)
// 	if _, err := rand.Read(nonce); err != nil {
// 		return err
// 	}

// 	ciphertext := gcm.Seal(nil, nonce, plaintext, nil)

// 	// Output: salt | nonce | ciphertext
// 	out := append(salt, nonce...)
// 	out = append(out, ciphertext...)

// 	return ioutil.WriteFile(outputPath, out, 0644)
// }

// // DecryptFile decrypts the input file and writes the result to output file.
// func DecryptFile(password, inputPath, outputPath string) error {
// 	data, err := ioutil.ReadFile(inputPath)
// 	if err != nil {
// 		return err
// 	}

// 	if len(data) < saltSize+nonceSize {
// 		return fmt.Errorf("file too short")
// 	}

// 	salt := data[:saltSize]
// 	nonce := data[saltSize : saltSize+nonceSize]
// 	ciphertext := data[saltSize+nonceSize:]

// 	key, err := DeriveKey([]byte(password), salt)
// 	if err != nil {
// 		return err
// 	}

// 	block, err := aes.NewCipher(key)
// 	if err != nil {
// 		return err
// 	}

// 	gcm, err := cipher.NewGCM(block)
// 	if err != nil {
// 		return err
// 	}

// 	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
// 	if err != nil {
// 		return err
// 	}

// 	return ioutil.WriteFile(outputPath, plaintext, 0644)
// }

// func main() {
// service := "potok"
// user := "my-vaults"
// var password string

// fmt.Scan(&password)

// err := keyring.Set(service, user, password)
// if err != nil {
// 	log.Fatal(err)
// }

// secret, err := keyring.Get(service, user)
// if err != nil {
// 	log.Fatal(err)
// }

// err = EncryptFile(secret, "plain.txt", "encrypted.txt")
// if err != nil {
// 	fmt.Println("Encryption error:", err)
// 	return
// }
// fmt.Println("File encrypted.")

// err = DecryptFile(secret, "encrypted.txt", "decrypted.txt")
// if err != nil {
// 	fmt.Println("Decryption error:", err)
// 	return
// }
// fmt.Println("File decrypted.")
// }

// func main() {
// 	watcher, err := fsnotify.NewWatcher()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer watcher.Close()

// 	// Recursively add all subdirectories
// 	err = filepath.Walk("/home/athena/me/Athena", func(path string, info os.FileInfo, err error) error {
// 		if err != nil {
// 			return err
// 		}
// 		if info.IsDir() {
// 			log.Println("Watching:", path)
// 			return watcher.Add(path)
// 		}
// 		return nil
// 	})
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	go func() {
// 		for {
// 			select {
// 			case event, ok := <-watcher.Events:
// 				if !ok {
// 					return
// 				}
// 				log.Println("event:", event)

// 				log.Println(event.Name)

// 				// If a new directory is created, add it to the watcher
// 				if event.Op&fsnotify.Create == fsnotify.Create {
// 					fi, err := os.Stat(event.Name)
// 					if err == nil && fi.IsDir() {
// 						log.Println("New directory detected, watching:", event.Name)
// 						watcher.Add(event.Name)
// 					}
// 				}
// 			case err, ok := <-watcher.Errors:
// 				if !ok {
// 					return
// 				}
// 				log.Println("error:", err)
// 			}
// 		}
// 	}()

// 	// Block forever
// 	select {}
// }
