package main

import "fmt"

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

func main() {
	totalDaSoma := soma(1, 2, 3, 4, 5, 6, 7, 34, 56, 200)
	fmt.Println(totalDaSoma)
}

// Com INT e STRING

// func escrever(texto string, numeros ...int) {
// 	for _, numero := range numeros {
// 		fmt.Println(texto, numero)
// 	}
// }

// func main() {
// 	totalDaSoma := soma(1, 2, 3, 4)
// 	fmt.Println(totalDaSoma)

// 	escrever("Olá Mundo", 1, 2, 3, 4, 5)
// }
