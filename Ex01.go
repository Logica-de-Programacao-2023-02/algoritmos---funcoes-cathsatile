package main

import "fmt"

func Media(slice []int) int {
	soma := 0
	for _, c := range slice {
		soma += c
	}
	media := soma / len(slice)
	return media
}
func main() {
	valores := []int{1, 2, 3, 4, 5}
	MediaResultado := Media(valores)
	fmt.Print("A média dos valores é: ", MediaResultado)
}
