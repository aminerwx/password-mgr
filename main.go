package main

import (
	"encoding/hex"
	"fmt"

	"github.com/aminerwx/password-mgr/core"
	"github.com/aminerwx/password-mgr/utils"
)

func main() {
	var pwd core.Password
	pwd.Length = 20
	pwd.HasUpper = true
	pwd.HasLower = true
	pwd.HasDigit = true
	pwd.HasSymbol = true
	pwd.Generate()
	fmt.Println(pwd.Text)

	// KDF options
	options := &core.Options{
		SaltLength:  32,
		KeyLength:   32,
		Iterations:  10,
		Memory:      128 * 1024,
		Parallelism: 2,
	}

	hash, err := core.CreateHash("password", options)
	utils.Maybe(err)

	match, _, err := core.VerifyHash("password", hash)
	utils.Maybe(err)

	k, _, _, err := core.DecodeHash(hash)
	utils.Maybe(err)

	if match {
		fmt.Println("Password is matching.")
		ciphertxt, err := core.EncryptAES([]byte("SecretMsg"), []byte(k))
		utils.Maybe(err)

		plaintext, err := core.DecryptAES(ciphertxt, []byte(k))
		utils.Maybe(err)

		cipher := hex.EncodeToString(ciphertxt)
		plain := string(plaintext)
		data := fmt.Sprintf("Secret Key: \n\t%v\nAES-256:\n\tciphertext: \n\t\t%v\n\tplaintext: \n\t\t%v\n", hash, cipher, plain)
		utils.WriteFile("./out/data", []byte(data))
	} else {
		fmt.Println("Incorrect password.")
	}
}
