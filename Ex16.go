package main

import (
	"errors"
	"fmt"
)

func substituirVogaisPorAsterisco(str string) (string, error) {
	if str == "" {
		return "", errors.New("A string está vazia")
	}

	strModificada := ""
	for _, char := range str {
		switch char {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			strModificada += "*"
		default:
			strModificada += string(char)
		}
	}

	return strModificada, nil
}

func main() {
	exemploString := "Hello, World! This is an example string."
	stringModificada, err := substituirVogaisPorAsterisco(exemploString)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("String modificada:", stringModificada)
	}
}
