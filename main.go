package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	password := generatePassword(8)
	fmt.Println(password)
}

func generatePassword(length int) string {
	lowerCase := "abcdefghijklmnopqrstuvwxyz"
	upperCase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers := "0123456789"
	special := "!@#$%^&*()+?><:{}[]"
	characters := lowerCase + upperCase + numbers + special

	var sb strings.Builder

	for i := 0; i < length; i++ {
		index := rand.Intn(len(characters))
		char := string(characters[index])
		sb.WriteString(char)
	}

	return sb.String()
}

// Priemiro momento -> gerar uma senha completamente aleatória
// Segundo momento -> gerar uma senha aleatória MAS com obrigatoriedade: letra maiuscula, número e caractere especial
