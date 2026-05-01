package main

import (
	"fmt"
	"math/rand"
)

func main() {
	password := generatePassword(12)
	fmt.Println(password)
}

func generatePassword(length int) string {
	lowerCase := "abcdefghijklmnopqrstuvwxyz"
	upperCase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers := "0123456789"
	special := "!@#$%^&*()+?><:{}[]"
	characters := lowerCase + upperCase + numbers + special

	// Adiciona os 4 caracteres obrigatórios
	chars := []byte{}
	chars = append(chars, lowerCase[rand.Intn(len(lowerCase))])
	chars = append(chars, upperCase[rand.Intn(len(upperCase))])
	chars = append(chars, numbers[rand.Intn(len(numbers))])
	chars = append(chars, special[rand.Intn(len(special))])

	//Adiciona o resto dos caracteres depois dos obrigatórios igual o tamanho inserido da senha
	for len(chars) < length {
		chars = append(chars, characters[rand.Intn(len(characters))])
	}

	//Embaralha as posições dos caracteres
	for i := len(chars) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		chars[i], chars[j] = chars[j], chars[i]
	}

	//Retorna a senha convertendo o tipo byte em string
	return string(chars)
}

// Priemiro momento -> gerar uma senha completamente aleatória
// Segundo momento -> gerar uma senha aleatória MAS com obrigatoriedade: letra maiuscula, número e caractere especial
