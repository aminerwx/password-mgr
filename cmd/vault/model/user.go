package model

import (
	"errors"

	"github.com/aminerwx/password-mgr/internal"
	"github.com/google/uuid"
)

type User struct {
	ID             string `json:"id"`
	Username       string `json:"username"`
	MasterPassword string `json:"master_password"`
}

func NewUser(username, masterPassword string) (User, error) {
	if len(username) == 0 {
		return User{}, errors.New("func NewUser: empty username")
	}
	hash, err := internal.NewHash(masterPassword, &internal.MyArgon2idOptions)
	if err != nil {
		return User{}, err
	}
	id := uuid.New()
	return User{ID: id.String(), Username: username, MasterPassword: hash}, nil
}
