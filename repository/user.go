package repository

import (
	"context"

	"github.com/aminerwx/password-mgr/model"
)

type UserRepository interface {
	GetUser(username string) (model.User, error)
	CreateUser(user model.User) error
	UpdateUser(newUser model.User) error
	RemoveUser(id string) error
}

func (p *PostgresStorage) GetUser(username string) (model.User, error) {
	stmt := `SELECT id, username, master_password FROM users WHERE username = $1;`
	var user model.User
	err := p.pool.QueryRow(context.Background(), stmt, username).Scan(&user.ID, &user.Username, &user.MasterPassword)
	if err != nil {
		return model.User{}, err
	}
	return user, err
}

func (p *PostgresStorage) CreateUser(user model.User) error {
	stmt := `INSERT INTO users(id, username, master_password) VALUES($1, $2, $3);`
	_, err := p.pool.Exec(context.Background(), stmt, user.ID, user.Username, user.MasterPassword)
	return err
}

func (p *PostgresStorage) UpdateUser(user model.User) error {
	stmt := `UPDATE users SET id = $1, username = $2, master_password = $3;`
	_, err := p.pool.Exec(context.Background(), stmt, user.ID, user.Username, user.MasterPassword)
	return err
}

func (p *PostgresStorage) RemoveUser(id string) error {
	stmt := `DELETE FROM users WHERE id = $1;`
	_, err := p.pool.Exec(context.Background(), stmt, id)
	return err
}
