package main

import "fmt"

func SegundoMaior(slice []int) int {
	var m, m2 int
	for i, s := range slice {
		if i == 0 || s > m {
			m2 = m
			m = s
		} else if i == 1 || s > m2 {
			m2 = s
		}
	}
	return m2
}
func main() {
	slice := []int{1, 3, 6, 45, 78, 81, 10}
	num := SegundoMaior(slice)
	fmt.Print("O segunfo maior número da slice é: ", num)
}
