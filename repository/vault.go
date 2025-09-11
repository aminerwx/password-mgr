package repository

import (
	"context"

	"github.com/aminerwx/password-mgr/model"
)

type VaultRepository interface {
	GetVault(username, master string) (model.Vault, error)
	CreateVault(user model.User) error
	UpdateVault(id string, newUser model.User, newMaster string) error
	RemoveVault(id string) error
}

func (p *PostgresStorage) GetVault(username, master string) (model.Vault, error) {
	return model.Vault{}, nil
}

func (p *PostgresStorage) CreateVault(user model.User) error {
	vault := model.NewVault(user.ID, user.MasterPassword)
	stmt := `INSERT INTO vault(user_id, encrypted_data) VALUES ($1, $2);`
	_, err := p.pool.Exec(context.Background(), stmt, vault.UserID, vault.Data)
	return err
}

func (p *PostgresStorage) UpdateVault(id string, newUser model.User, newMaster string) error {
	return nil
}

func (p *PostgresStorage) RemoveVault(id string) error {
	return nil
}
