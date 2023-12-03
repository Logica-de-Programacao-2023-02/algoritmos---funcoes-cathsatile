package main

import (
	"errors"
	"fmt"
	"sort"
)

func Ordernar(slice []int) ([]int, error) {
	if len(slice) == 0 {
		return nil, errors.New(" Slice está vazio")
	}
	copiaOrdenada := make([]int, len(slice))
	copy(copiaOrdenada, slice)

	sort.Ints(copiaOrdenada)

	return copiaOrdenada, nil
}

func main() {
	exemplo := []int{5, 2, 8, 1, 7}
	slice, err := Ordernar(exemplo)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Valores ordenados de forma crescente:", slice)
	}
}
