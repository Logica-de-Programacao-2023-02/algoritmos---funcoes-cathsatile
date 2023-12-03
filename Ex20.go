package main

import (
	"errors"
	"fmt"
)

func stringsComMaisDeCincoCaracteres(sliceStrings []string) ([]string, error) {
	if len(sliceStrings) == 0 {
		return nil, errors.New("O slice de strings está vazio")
	}

	var resultado []string
	for _, str := range sliceStrings {
		if len(str) > 5 {
			resultado = append(resultado, str)
		}
	}

	return resultado, nil
}

func main() {
	sliceExemplo := []string{"apple", "banana", "orange", "grape", "kiwi"}
	resultado, err := stringsComMaisDeCincoCaracteres(sliceExemplo)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Strings com mais de 5 caracteres:", resultado)
	}
}
