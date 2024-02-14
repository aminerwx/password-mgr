package main

import (
	crand "crypto/rand"
	"fmt"
	"math/big"
	mrand "math/rand"
	"strings"
)

type Password struct {
	Text      string
	Length    int
	Entropy   float32
	HasUpper  bool
	HasLower  bool
	HasDigit  bool
	HasSymbol bool
}

func main() {
	var pwd Password
	pwd.Length = 15
	pwd.HasUpper = true
	pwd.HasLower = true
	pwd.HasDigit = true

	pwd.Generate()
}

func (p *Password) Generate() {
	var password strings.Builder
	counter := 0
	if p.HasUpper {
		counter += 1
		password.WriteString(GenerateUpper(5))
	}
	if p.HasLower {
		counter += 1
		password.WriteString(GenerateLower(5))
	}
	if p.HasDigit {
		counter += 1
		password.WriteString(GenerateDigit(5))
	}
	shuff := []rune(password.String())
	mrand.Shuffle(len(shuff), func(i, j int) {
		shuff[i], shuff[j] = shuff[j], shuff[i]
	})
	fmt.Printf("Unshuffled: \t%v\n", password.String())
	p.Text = string(shuff)
	fmt.Printf("Shuffled: \t%v\n", p.Text)
}

func (p *Password) String() string { return p.Text }

func GenerateUpper(length int) string {
	var sb strings.Builder
	for i := 0; i < length; i++ {
		uppers := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		index, _ := crand.Int(crand.Reader, big.NewInt(int64(26)))
		sb.WriteString(string(uppers[index.Int64()]))
	}
	return sb.String()
}

func GenerateLower(length int) string {
	var sb strings.Builder
	for i := 0; i < length; i++ {
		uppers := "abcdefghijklmnopqrstuvwxyz"
		index, _ := crand.Int(crand.Reader, big.NewInt(int64(26)))
		sb.WriteString(string(uppers[index.Int64()]))
	}
	return sb.String()
}

func GenerateDigit(length int) string {
	var sb strings.Builder
	for i := 0; i < length; i++ {
		uppers := "0123456789"
		index, _ := crand.Int(crand.Reader, big.NewInt(int64(10)))
		sb.WriteString(string(uppers[index.Int64()]))
	}
	return sb.String()
}

/*
func GenerateSymbol(length int) string {
	var sb strings.Builder
	for i := 0; i < length; i++ {
		uppers := "!\"#$%&\'()*+,-./:;<=>?@[\\]^_\`{|}~"
		index, _ := rand.Int(rand.Reader, big.NewInt(int64(26)))
		sb.WriteString(string(uppers[index.Int64()]))
	}
	return sb.String()
}
*/
