package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/aminerwx/password-mgr/cmd/vault/api"
	"github.com/aminerwx/password-mgr/cmd/vault/model"
	"github.com/aminerwx/password-mgr/cmd/vault/repository"
	"github.com/aminerwx/password-mgr/core"
	"github.com/jackc/pgx/v5/pgxpool"
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
			Password: securePwd.NewRandomString(),
		},
		{
			Title:    "amazon",
			Username: "user@amazon.com",
			Password: securePwd.NewRandomString(),
		},
		{
			Title:    "paypal",
			Username: "user@paypal.com",
			Password: securePwd.NewRandomString(),
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

	key, _, _, err := core.Decode(dorakyura.MasterPassword)
	encryptedData, err := core.Encrypt(data, key)
	// TODO:
	// Save data to database
	// Save encryptedData to file
	fmt.Println("key =", string(key))
	v := model.NewVault(dorakyura.ID, dorakyura.MasterPassword)
	v.Save(myroot)
	fmt.Printf("%+v\n", v)
	os.WriteFile("root.json", encryptedData, 0644)
	fmt.Println("Server started at localhost:3333")
	uri := "postgresql://postgres@localhost:5432/password_mgr?sslmode=disable"
	pool, err := pgxpool.New(context.Background(), uri)
	defer pool.Close()
	if err != nil {
		log.Fatalf("unable to connect to database: %v", err)
	}
	store := repository.NewPostgresRepository(pool)

	srv := api.NewServer(store, ":3333")
	if err := srv.Start(); err != nil {
		panic(err)
	}
}

// Encrypt Vault
// CRUD Vault
// CRUD User
