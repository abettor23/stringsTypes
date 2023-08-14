package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Исходная строка:")

	smain := `a10 10 20b 20 30c30 30 dd`

	fmt.Println(smain)

	var numbers []string

	elements := strings.Split(smain, " ")

	for i := 0; i < len(elements); i++ {

		currentId := elements[i]

		check, err := strconv.Atoi(currentId)

		if err == nil {
			numbers = append(numbers, strconv.Itoa(check))
		}
	}

	fmt.Println("В строке содержатся числа в десятичном формате:")
	fmt.Println(strings.Join(numbers, " "))
}
