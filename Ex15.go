package main

import (
	"errors"
	"fmt"
)

func NumeroNoSlice(numero int, slice []int) (bool, error) {
	if len(slice) == 0 {
		return false, errors.New(" Slice  está vazio")
	}

	for _, valor := range slice {
		if valor == numero {
			return true, nil
		}
	}

	return false, nil
}

func main() {
	numexemplo := 3
	slicexemplo := []int{1, 2, 3, 4, 5}

	NoSlice, err := NumeroNoSlice(numexemplo, slicexemplo)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("O número %d está no slice? %t\n", numexemplo, NoSlice)
	}
}
