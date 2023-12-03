package main

import (
	"errors"
	"fmt"
	"strings"
)

func ConcatenarVirgulas(slice []string) (string, error) {

	if len(slice) == 0 {
		return "", errors.New("Slice está vazio")
	}

	resultado := strings.Join(slice, ",")
	return resultado, nil
}

func main() {
	slice := []string{"Onça ", "é ", "um", " animal muito confundido com Jaguar."}
	resultado, err := ConcatenarVirgulas(slice)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Concatenação com vírgulas:", resultado)
	}

}
