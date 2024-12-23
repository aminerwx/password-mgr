package core

import (
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"github.com/aminerwx/password-mgr/utils"
	"golang.org/x/crypto/argon2"
)

type Argon2idOptions struct {
	Iterations  uint32
	Memory      uint32
	Parallelism uint8
}

var MyArgon2idOptions = NewArgon2idOptions(10, 50, 12)

func NewArgon2idOptions(iterations, mb uint32, parallel uint8) Argon2idOptions {
	return Argon2idOptions{
		Iterations:  iterations,
		Memory:      mb * 1024,
		Parallelism: parallel,
	}
}

// Generate secret key and salt (argon2id KDF) from passphrase and options.
func GenerateKey(passphrase string, options *Argon2idOptions) (key, salt []byte, err error) {
	pwd := []byte(passphrase)
	if len(pwd) == 0 {
		return nil, nil, errors.New("empty passphrase")
	}

	salt, err = utils.RandomBytes(32)
	if err != nil {
		return nil, nil, err
	}

	key = argon2.IDKey(pwd, salt, options.Iterations, options.Memory, options.Parallelism, 32)
	return key, salt, nil
}

// Create an argon2id string representation from secret key, salt and options
func CreateHash(passphrase string, options *Argon2idOptions) (hash string, err error) {
	key, salt, err := GenerateKey(passphrase, options)
	if err != nil {
		return "", err
	}

	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Key := base64.RawStdEncoding.EncodeToString(key)

	hash = fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, options.Memory, options.Iterations, options.Parallelism, b64Salt, b64Key)
	return hash, nil
}

// Extract secret key, salt and options from argon2id string representation
func DecodeHash(hash string) (key, salt []byte, options *Argon2idOptions, err error) {
	values := strings.Split(hash, "$")
	if len(values) != 6 {
		return nil, nil, nil, errors.New("argon2id format is invalid")
	}
	if values[1] != "argon2id" {
		return nil, nil, nil, errors.New("invalid argon2 variant")
	}
	var version int
	_, err = fmt.Sscanf(values[2], "v=%d", &version)
	if err != nil {
		return nil, nil, nil, err
	}
	if argon2.Version != version {
		return nil, nil, nil, errors.New("invalid Argon2 version")
	}
	options = &Argon2idOptions{}
	fmt.Sscanf(values[3], "m=%d,t=%d,p=%d", &options.Memory, &options.Iterations, &options.Parallelism)

	salt, err = base64.RawStdEncoding.Strict().DecodeString(values[4])
	if err != nil {
		return nil, nil, nil, err
	}

	key, err = base64.RawStdEncoding.Strict().DecodeString(values[5])
	if err != nil {
		return nil, nil, nil, err
	}

	return key, salt, options, nil
}

// Compare passphrase and decoded hash
func VerifyHash(passphrase string, hash string) (match bool, options *Argon2idOptions, err error) {
	key, salt, o, err := DecodeHash(hash)
	options = o
	if err != nil {
		return match, options, err
	}
	otherKey := argon2.IDKey([]byte(passphrase), salt, o.Iterations, o.Memory, o.Parallelism, 32)

	keyLen := int32(len(key))
	otherKeyLen := int32(len(otherKey))
	if subtle.ConstantTimeEq(keyLen, otherKeyLen) == 0 {
		return false, options, nil
	}
	if subtle.ConstantTimeCompare(key, otherKey) == 1 {
		return true, options, nil
	}

	return match, options, err
}
