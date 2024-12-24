package model

import (
	"github.com/aminerwx/password-mgr/core"
	"github.com/google/uuid"
)

type User struct {
	ID             string `json:"id"`
	Username       string `json:"username"`
	MasterPassword string `json:"master_password"`
}

func NewUser(username, masterPassword string) (User, error) {
	hash, err := core.NewHash(masterPassword, &core.MyArgon2idOptions)
	if err != nil {
		return User{}, err
	}
	id := uuid.New()
	return User{ID: id.String(), Username: username, MasterPassword: hash}, nil
}
