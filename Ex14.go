package main

import (
	"errors"
	"fmt"
)

func IntersecaoSlices(slice1, slice2 []int) ([]int, error) {
	if len(slice1) == 0 || len(slice2) == 0 {
		return nil, errors.New(" Algum slice está vazio")
	}

	mapaSlice1 := make(map[int]bool)
	for _, num := range slice1 {
		mapaSlice1[num] = true
	}

	intersecao := make([]int, 0)

	for _, num := range slice2 {
		if mapaSlice1[num] {
			intersecao = append(intersecao, num)
		}
	}

	return intersecao, nil
}

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{3, 4, 5, 6, 7}

	resultado, err := IntersecaoSlices(slice1, slice2)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Interseção dos slices:", resultado)
	}
}
