package main

func main() {
	var pwd Password
	pwd.Length = 15
	pwd.HasUpper = true
	pwd.HasLower = true
	pwd.HasDigit = true
	pwd.Generate()
}
