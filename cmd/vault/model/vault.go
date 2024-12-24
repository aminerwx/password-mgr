package model

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/aminerwx/password-mgr/core"
	"github.com/google/uuid"
)

type Record struct {
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	AccessedAt time.Time `json:"accessed_at"`
}

func NewRecord() Record {
	now := time.Now()
	return Record{
		CreatedAt:  now,
		ModifiedAt: now,
		AccessedAt: now,
	}
}

func (r *Record) UpdateAccessedAt() {
	r.AccessedAt = time.Now()
}

func (r *Record) UpdateModifiedAt() {
	r.ModifiedAt = time.Now()
}

type Credential struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Url       string `json:"url"`
	Tag       string `json:"tag"`
	Timestamp Record `json:"timestamp"`
}

func NewCredential(title, username, password string) Credential {
	return Credential{
		Title:     title,
		Username:  username,
		Password:  password,
		Timestamp: NewRecord(),
	}
}

type Group struct {
	Title     string       `json:"title"`
	Entries   []Credential `json:"entries"`
	SubGroups []Group      `json:"subgroups"`
}

func NewGroup(title string) Group {
	return Group{Title: title}
}

func (g Group) AddEntry(entry Credential) {
	g.Entries = append(g.Entries, entry)
}

func (g Group) RemoveEntry(id string) {
	for k, v := range g.Entries {
		if id == v.ID {
			if k == len(g.Entries)-1 {
				g.Entries = g.Entries[:k]
				break
			} else {
				g.Entries = append(g.Entries[:k], g.Entries[:k+1]...)
				break
			}
		}
	}
}

type Vault struct {
	ID                 string `json:"id"`
	UserID             string `json:"user_id"`
	MasterPasswordHash string `json:"master_password_hash"`
	EncryptedData      []byte `json:"encrypted_data"`
	Timestamp          Record `json:"timestamp"`
}

func NewVault(userID string, master string) Vault {
	id := uuid.New()
	return Vault{
		ID:                 id.String(),
		Timestamp:          NewRecord(),
		UserID:             userID,
		MasterPasswordHash: master,
	}
}

func (v Vault) Save(root Group) error {
	_, err := json.Marshal(root)
	if err != nil {
		log.Fatalln("error marshal")
	}
	opts := core.NewArgon2idOptions(10, 20, 2)
	hash, err := core.NewHash("pwd", &opts)
	if err != nil {
		return err
	}
	fmt.Println(hash)
	return nil
}
