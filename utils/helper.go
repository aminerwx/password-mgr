package utils

import (
	"crypto/rand"
	"fmt"
	"os"
)

func RandomBytes(n uint32) ([]byte, error) {
	buf := make([]byte, n)
	if _, err := rand.Read(buf); err != nil {
		return nil, err
	}
	return buf, nil
}

func ReadFile(filepath string) ([]byte, error) {
	dat, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	return dat, nil
}

func WriteFile(filepath string, data []byte) error {
	fmt.Println("File written.")
	if err := os.WriteFile(filepath, data, 0600); err != nil {
		return err
	}
	return nil
}

func Maybe(err error) {
	if err != nil {
		panic(err)
	}
}
