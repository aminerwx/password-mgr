package main

import (
	"fmt"

	"github.com/aminerwx/password-mgr/core"
	"github.com/aminerwx/password-mgr/utils"
)

// TODO:
// Create object structure
// Serialize it and encrypt it
type Credential struct {
	Title    string        `json:"title"`
	Username string        `json:"username"`
	Password core.Password `json:"password"`
	Url      string        `json:"url"`
}

type File struct {
	MasterPassword []byte
	Data           []Credential
}

func main() {
	var pwd core.Password
	pwd.Length = 20
	pwd.HasUpper = true
	pwd.HasLower = true
	pwd.HasDigit = true
	pwd.HasSymbol = true
	pwd.Generate()
	fmt.Println(pwd.String())

	// store master password hash in db
	hash, err := core.CreateHash("master password", &core.MyArgon2idOptions)
	utils.Maybe(err)
	fmt.Println(hash)

	// retrieve hash and compare
	// match, _, err := core.VerifyHash("master password", hash)
	// utils.Maybe(err)

	// if match {
	//	fmt.Println("Password is matching.")
	//	k, _, _, err := core.DecodeHash(hash)
	//	fmt.Println(k)
	//	utils.Maybe(err)

	//	// Encrypt User Data
	//	cipher, err := core.EncryptAES([]byte("SecretMsg"), []byte(k))
	//	utils.Maybe(err)

	//	// Decrypt User Data
	//	plaintext, err := core.DecryptAES(cipher, []byte(k))
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
