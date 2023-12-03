package main

import "fmt"

func Concatenacao(s []string) string {
	resultado := ""
	for _, c := range s {
		resultado += c
	}
	return resultado
}

func main() {
	slice := []string{"Onça ", "é ", "um", " animal muito confundido com Jaguar."}
	resultado := Concatenacao(slice)
	fmt.Println("A concatenação das strings é:", resultado)

}
