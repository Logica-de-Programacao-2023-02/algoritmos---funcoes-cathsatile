package main

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

func ordenarStrings(sliceStrings []string) (string, error) {
	if len(sliceStrings) == 0 {
		return "", errors.New("O slice de strings está vazio")
	}

	sort.Strings(sliceStrings)
	strOrdenada := strings.Join(sliceStrings, " ")

	return strOrdenada, nil
}

func main() {
	sliceExemplo := []string{"banana", "abacaxi", "laranja", "uva"}
	stringOrdenada, err := ordenarStrings(sliceExemplo)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Strings ordenadas:", stringOrdenada)
	}
}
