package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/aminerwx/password-mgr/cmd/vault/model"
	"github.com/aminerwx/password-mgr/core"
)

// Fill Entry
func main() {
	securePwd := core.Password{
		HasUpper:  true,
		HasLower:  true,
		HasDigit:  true,
		HasSymbol: true,
	}

	dummiesCredential := []model.Credential{
		{
			Title:    "Arch",
			Username: "amine",
			Password: "qwerty",
		},
		{
			Title:    "gmail",
			Username: "user@gmail.com",
			Password: securePwd.Generate(),
		},
		{
			Title:    "amazon",
			Username: "user@amazon.com",
			Password: securePwd.Generate(),
		},
		{
			Title:    "paypal",
			Username: "user@paypal.com",
			Password: securePwd.Generate(),
		},
	}

	online := model.Group{
		Title:   "Online",
		Entries: dummiesCredential,
	}

	myroot := model.Group{
		Title: "Root",
		Entries: []model.Credential{
			dummiesCredential[0],
		},
		SubGroups: []model.Group{
			online,
		},
	}

	// fmt.Println(myroot.Title)
	// fmt.Println(myroot.Entries)
	// fmt.Println(myroot.SubGroups)

	data, err := json.MarshalIndent(myroot, "", "  ")
	if err != nil {
		log.Fatalln("error marshall")
	}

	dorakyura, err := model.NewUser("dorakyura", "qwerty")
	if err != nil {
		panic(err)
	}

	key, _, _, err := core.DecodeHash(dorakyura.MasterPasswordHash)
	encryptedData, err := core.EncryptAES(data, key)
	// TODO:
	// Save data to database
	// Save encryptedData to file
	v := model.NewVault(dorakyura.ID, dorakyura.MasterPasswordHash)
	v.Save(myroot)
	fmt.Printf("%+v\n", v)
	os.WriteFile("root.json", encryptedData, 0644)
}

// Encrypt Vault
// CRUD Vault
// CRUD User
