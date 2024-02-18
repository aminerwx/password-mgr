package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"

	"golang.org/x/crypto/hkdf"
)

func Encrypt(masterPassword []byte) {
	hash := sha256.New

	salt := make([]byte, hash().Size())
	if _, err := rand.Read(salt); err != nil {
		panic(err)
	}
	info := []byte("")
	//	data := []byte("My Data.")
	h := hkdf.New(hash, masterPassword, salt, info)

	// secretHex := hex.EncodeToString(masterPassword)

	fmt.Println(h, hash().Size())
}

func RandomBytes(n uint32) ([]byte, error) {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return nil, err
	}
	return b, nil
}

/*
 *  Encrypt:
 *
 *    - Input Binary -> StringDecode ------
 *
 *                                        > AES(DecodedBin, )
 *    HKDF inputs:
 *      Hash function i.e sha256
 *      Source key = key from which multiple keys can be derived
 *      length = number of bytes to derive
 *      contextInfo = arbitrary string used to bind a derived key to an intended context
 *      salt = optional extra randomness (recommended hash-length random value)
 *
 *
 *
 * */
