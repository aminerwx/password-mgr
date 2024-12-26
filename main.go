package main

import (
	"fmt"

	"github.com/aminerwx/password-mgr/internal"
	"github.com/aminerwx/password-mgr/utils"
)

// TODO:
// Create object structure
// Serialize it and encrypt it
type Credential struct {
	Title    string            `json:"title"`
	Username string            `json:"username"`
	Password internal.Password `json:"password"`
	Url      string            `json:"url"`
}

type File struct {
	MasterPassword []byte
	Data           []Credential
}

func main() {
	var pwd internal.Password
	pwd.Length = 20
	pwd.HasUpper = true
	pwd.HasLower = true
	pwd.HasDigit = true
	pwd.HasSymbol = true
	pwd.NewRandomString()
	fmt.Println(pwd.ToString())

	// store master password hash in db
	hash, err := internal.NewHash("master password", &internal.MyArgon2idOptions)
	utils.Maybe(err)
	fmt.Println(hash)

	// retrieve hash and compare
	// match, _, err := internal.VerifyHash("master password", hash)
	// utils.Maybe(err)

	// if match {
	//	fmt.Println("Password is matching.")
	//	k, _, _, err := internal.DecodeHash(hash)
	//	fmt.Println(k)
	//	utils.Maybe(err)

	//	// Encrypt User Data
	//	cipher, err := internal.EncryptAES([]byte("SecretMsg"), []byte(k))
	//	utils.Maybe(err)

	//	// Decrypt User Data
	//	plaintext, err := internal.DecryptAES(cipher, []byte(k))
	//	utils.Maybe(err)

	//	ciphertxt := hex.EncodeToString(cipher)
	//	plain := string(plaintext)
	//	data := fmt.Sprintf("Secret Key: \n\t%v\nAES-256:\n\tciphertext: \n\t\t%v\n\tplaintext: \n\t\t%v\n", hash, ciphertxt, plain)

	//	pwd, _ := os.Getwd()
	//	out := "data.txt"

	//	utils.WriteFile(path.Join(pwd, "cipher"), cipher)

	//	utils.WriteFile(path.Join(pwd, out), []byte(data))
	//} else {
	//	fmt.Println("Incorrect password.")
	//}
}
