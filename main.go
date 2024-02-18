package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"github.com/aminerwx/password-mgr/core"
)

/*
* KDF -> Derive Key from passphrase and feed it to AES
* AES -> Encrypt data
* */

func main() {
	/*
		var pwd password.Password
		pwd.Length = 15
		pwd.HasUpper = true
		pwd.HasLower = true
		pwd.HasDigit = true
		fmt.Println(pwd.Generate())
		uppers := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		str := password.GenerateUpper(5)
		fmt.Println(strings.Contains(uppers, str))
		//	utils.Encrypt([]byte("secret"))
	*/
	plaintext := []byte("Secret Message")
	password := []byte("MyPassword123456")
	fmt.Println(Encrypt(plaintext, password))
	options := &core.Options{
		SaltLength:  32,
		KeyLength:   32,
		Iterations:  4,
		Memory:      256 * 1024,
		Parallelism: 6,
	}
	hash, err := core.CreateHash("password", options)
	if err != nil {
		panic(err)
	}
	fmt.Println("Hash: ", hash)
	k, s, o, err := core.DecodeHash(hash)
	if err != nil {
		panic(err)
	}
	fmt.Println("Decoded Hash: ", string(k), string(s), o)
	match, _, err := core.VerifyHash("password", hash)
	if err != nil {
		panic(err)
	}
	fmt.Println(match)
}

// TODO: AES256 encryption
func Encrypt(plaintext []byte, key []byte) string {
	aesBlock, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	if len(plaintext) == 0 {
		panic("plaintext is empty")
	}

	iv := make([]byte, aesBlock.BlockSize())
	if _, err := rand.Read(iv); err != nil {
		panic(err)
	}

	cbc := cipher.NewCBCEncrypter(aesBlock, iv)
	content := PKCS5Padding(plaintext, aesBlock.BlockSize())
	ciphertext := make([]byte, len(content))
	cbc.CryptBlocks(ciphertext, content)
	return hex.EncodeToString(ciphertext)
}

func Decrypt(ciphertext []byte, key []byte) string {
	aesBlock, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	if len(ciphertext) == 0 {
		panic("ciphertext is empty")
	}

	iv := make([]byte, aesBlock.BlockSize())
	if _, err := rand.Read(iv); err != nil {
		panic(err)
	}
	return ""
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
