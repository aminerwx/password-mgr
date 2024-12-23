package repository

import "github.com/aminerwx/password-mgr/cmd/vault/model"

type VaultRepository interface {
	GetVault(id string) (model.Vault, error)
	CreateVault(user model.User, master string) error
	UpdateVault(id string, newUser model.User, newMaster string) error
	RemoveVault(id string) error
}

func (p *PostgresStorage) GetVault(id string) (model.Vault, error) {
	return model.Vault{}, nil
}

func (p *PostgresStorage) CreateVault(user model.User, master string) error {
	return nil
}

func (p *PostgresStorage) UpdateVault(id string) error {
	return nil
}

func (p *PostgresStorage) RemoveVault(id string) error {
	return nil
}
