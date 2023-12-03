package main

import (
	"errors"
	"fmt"
)

func numerosPrimosAteN(numero int) ([]int, error) {
	if numero < 0 {
		return nil, errors.New("O número é negativo")
	}

	var primos []int
	for i := 2; i <= numero; i++ {
		if ehPrimo(i) {
			primos = append(primos, i)
		}
	}

	return primos, nil
}
func ehPrimo(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	numeroExemplo := 20
	primos, err := numerosPrimosAteN(numeroExemplo)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Números primos até %d: %v\n", numeroExemplo, primos)
	}
}
