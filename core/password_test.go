package core

import (
	"testing"
)

type Test struct {
	name   string
	got    int
	expect int
}

func TestGenerateUpper(t *testing.T) {
	//	test := Test{"5 should be 5", 5, 5}
	uppers := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	str := GenerateUpper(5)
	acc := 0
	for _, s := range str {
		found := false
		for _, u := range uppers {
			if u == s {
				acc += 1
				found = true
				break
			}
		}
		if !found {
			break
		}
	}
	if len(str) != acc {
		t.Fatal("GenerateUpper doesn't return Uppercased-alphabet")
	}
	res := GenerateUpper(0)
	if len(res) != 0 {
		t.Fatalf("expected %v got %v\n", 0, len(res))
	}
}
