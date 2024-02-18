package core

import (
	crand "crypto/rand"
	"fmt"
	"math"
	"math/big"
	mrand "math/rand"
	"strings"
)

type Password struct {
	Text      string
	Length    int
	Charset   int
	Entropy   float64
	HasUpper  bool
	HasLower  bool
	HasDigit  bool
	HasSymbol bool
}

func (p *Password) Generate() string {
	var password strings.Builder
	if p.HasUpper {
		p.Charset += 26
		password.WriteString(GenerateUpper(5))
	}
	if p.HasLower {
		p.Charset += 26
		password.WriteString(GenerateLower(5))
	}
	if p.HasDigit {
		p.Charset += 10
		password.WriteString(GenerateDigit(5))
	}
	entropy := math.Log2(math.Pow(float64(p.Charset), float64(p.Length)))
	p.Entropy = math.Round(entropy*100) / 100
	shuff := []rune(password.String())
	mrand.Shuffle(len(shuff), func(i, j int) {
		shuff[i], shuff[j] = shuff[j], shuff[i]
	})
	p.Text = string(shuff)
	return p.String()
}

func (p *Password) String() string {
	var sb strings.Builder
	fmt.Fprintf(&sb, "Password = %v\tCharset = %v\tEntropy = %v bits", p.Text, p.Charset, p.Entropy)
	return sb.String()
}

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
		lowers := "abcdefghijklmnopqrstuvwxyz"
		index, _ := crand.Int(crand.Reader, big.NewInt(int64(26)))
		sb.WriteString(string(lowers[index.Int64()]))
	}
	return sb.String()
}

func GenerateDigit(length int) string {
	var sb strings.Builder
	for i := 0; i < length; i++ {
		numbers := "0123456789"
		index, _ := crand.Int(crand.Reader, big.NewInt(int64(10)))
		sb.WriteString(string(numbers[index.Int64()]))
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
