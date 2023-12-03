package main

import (
	"errors"
	"fmt"
)

func StringsMaiuscula(sliceStrings []string) (string, error) {
	if len(sliceStrings) == 0 {
		return "", errors.New(" Slice está vazio")
	}

	var result string
	for _, str := range sliceStrings {
		if len(str) > 0 && str[0] >= 'A' && str[0] <= 'Z' {
			result += str + " "
		}
	}

	return result[:len(result)-1], nil
}

func main() {
	exemplo := []string{"Gato", "cachorro", "pássaro", "Elefante", "onça", "Cavalo"}
	resultado, err := StringsMaiuscula(exemplo)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Strings com letra maiúscula: ", resultado)
	}
}
