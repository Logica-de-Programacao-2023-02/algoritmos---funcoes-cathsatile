package main

import (
	"fmt"
	"strings"
)

func QuantidadeVogais(s string) int {
	s = strings.ToLower(s)
	Vogais := 0

	for _, c := range s {
		switch c {
		case 'a', 'e', 'é', 'i', 'o', 'u':
			Vogais++
		}
	}
	return Vogais
}
func main() {
	x := "Hamburguer com picles e bacon é muito bom"
	Vogais := QuantidadeVogais(x)
	fmt.Printf("A quantidade de vogais na string é: %d\n", Vogais)
}
