package main

import (
	"errors"
	"fmt"
	"strings"
)

func Extrair(s string) ([]string, error) {
	if s == "" {
		return nil, errors.New("A string está vazia")
	}

	palavras := strings.Fields(s)
	return palavras, nil
}

func main() {
	exemplo := "Esta é uma string de exemplo"
	palavras, err := Extrair(exemplo)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Palavras na string:", palavras)
	}
}
