package main

import (
	"errors"
	"fmt"
)

func aplicarFuncao(slice []int, funcao func(int) int) ([]int, error) {
	if len(slice) == 0 {
		return nil, errors.New(" Slice está vazio")
	}

	resultados := make([]int, len(slice))
	for i, valor := range slice {
		resultados[i] = funcao(valor)
	}

	return resultados, nil
}
func main() {
	funcaoQuadrado := func(x int) int {
		return x * x
	}

	exemplo := []int{1, 2, 3, 4, 5}
	resultados, err := aplicarFuncao(exemplo, funcaoQuadrado)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Resultado da função:", resultados)
	}
}
