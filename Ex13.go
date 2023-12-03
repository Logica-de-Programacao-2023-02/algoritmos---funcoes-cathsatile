package main

import (
	"errors"
	"fmt"
)

func SomaDigitos(num int) (int, error) {
	if num < 0 {
		return 0, errors.New("O número é negativo")
	}

	soma := 0
	for num > 0 {
		digito := num % 10
		soma += digito
		num /= 10
	}

	return soma, nil
}

func main() {
	exemplo := 56789
	soma, err := SomaDigitos(exemplo)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("A soma dos dígitos de %d é: %d\n", exemplo, soma)
	}
}
