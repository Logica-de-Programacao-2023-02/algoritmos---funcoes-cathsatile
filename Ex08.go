package main

import (
	"errors"
	"fmt"
)

func Pares(slice []int) ([]int, error) {
	if len(slice) == 0 {
		return nil, errors.New(" Slice está vazio")
	}

	var pares []int
	for _, valor := range slice {
		if valor%2 == 0 {
			pares = append(pares, valor)
		}
	}

	return pares, nil
}

func main() {
	exemplo := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	resultado, err := Pares(exemplo)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Números pares no slice:", resultado)
	}
}
