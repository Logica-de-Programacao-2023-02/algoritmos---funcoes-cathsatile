package main

import "fmt"

func AcharPosicao(slice []int, valorP int) int {
	for i, valor := range slice {
		if valor == valorP {
			return i
		}
	}
	return -1
}
func main() {
	slice := []int{5, 67, 4, 31, 10, 28, 3}
	valor := 3
	posiçao := AcharPosicao(slice, valor)
	fmt.Printf("Esse valor está na posição índice=%d da slice.", posiçao)

}
