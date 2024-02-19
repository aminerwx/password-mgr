package utils

import (
	"crypto/rand"
)

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
