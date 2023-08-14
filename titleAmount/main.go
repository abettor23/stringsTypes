package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println("Определение количества слов, начинающихся с большой буквы в строке:")

	smain := `Go is an Open source programming Language that makes it Easy 
to build simple, reliable, and efficient Software`

	fmt.Println(smain)

	count := 0

	elements := strings.Split(smain, " ")

	for i := 0; i < len(elements); i++ {

		currentWord := elements[i]
		lowerCheck := strings.ToLower(currentWord)

		if currentWord != lowerCheck {
			count++
		}
	}

	fmt.Println("Строка содержит", count, "слов с большой буквы.")
}
