package main

import "fmt"

func main() {
	// ARITMETICOS
	soma := 1 + 2
	subtracao := 2 - 1
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	var numero1 int16 = 10
	var numero2 int16 = 25 // o GO só permite calculos com numeros de INT igual
	soma2 := numero1 + numero2
	fmt.Println(soma2)

	// FIM ARITMETICOS

	// ATRIBUIÇÂO

	var variavel1 string = "String1"
	variavel2 := "String2"
	fmt.Println(variavel1, variavel2)

	// FIM OPERADORES DE ATRIBUICAO

	// OPERADORES RELACIONAIS
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)

	// FIM DOS OPERADORES RELACIONAIS

	// OPERADORES LÓGICOS
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso) // E
	fmt.Println(verdadeiro || falso) // OU
	fmt.Println(!verdadeiro)         // NEGAÇÂO
	fmt.Println(!falso)

	//FIM DOS OPERADORES LÓGICOS

	//OPERADORES UNÁRIOS
	numero := 10
	numero++
	numero += 15 // numero = numero + 15
	fmt.Println(numero)

	numero--
	numero -= 15 // numero = numero - 15

	numero *= 3 // numero = numero * 3
	numero /= 10
	numero %= 3
	fmt.Println(numero)

	//FIM DOS OPERADORES UNÁRIOS

	// OPERADORES TERNÁRIOS

	var texto string
	if numero > 5 {
		texto = "Maior que 5"
	} else {
		texto = "Menor que 5"
	}

	fmt.Println(texto)

}
