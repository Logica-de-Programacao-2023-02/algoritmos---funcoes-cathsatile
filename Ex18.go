package main

import (
	"errors"
	"fmt"
)

func aplicarEsomar(sliceInteiros []int, funcao func(int) int) (int, error) {
	if len(sliceInteiros) == 0 {
		return 0, errors.New("O slice de inteiros está vazio")
	}

	soma := 0
	for _, valor := range sliceInteiros {
		soma += funcao(valor)
	}

	return soma, nil
}

func main() {
	funcaoQuadrado := func(x int) int {
		return x * x
	}

	sliceExemplo := []int{1, 2, 3, 4, 5}
	resultadoSoma, err := aplicarEsomar(sliceExemplo, funcaoQuadrado)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("A soma dos resultados é: %d\n", resultadoSoma)
	}
}
