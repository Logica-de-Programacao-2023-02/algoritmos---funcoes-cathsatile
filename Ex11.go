package main

import (
	"errors"
	"fmt"
)

func Primo(num int) (bool, error) {
	if num < 0 {
		return false, errors.New("O número é negativo")
	}
	if num < 2 {
		return false, nil
	}

	limite := num / 2
	for i := 2; i < limite; i++ {
		if num%i == 0 {
			return false, nil
		}
	}

	return true, nil
}

func main() {
	exemplo := 21
	primo, err := Primo(exemplo)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%d é primo? %t\n", exemplo, primo)
	}
}
