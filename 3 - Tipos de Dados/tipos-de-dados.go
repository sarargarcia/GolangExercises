package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int64 = 10000000
	fmt.Println(numero)

	var numero2 uint = 1000
	fmt.Println(numero2)

	//alias
	//INT32 = RUNE
	var numero3 rune = 1234
	fmt.Println(numero3)

	//BYTE = UINT8
	var numero4 byte = 123
	fmt.Println(numero4)

	//FIM NÚMEROS INTEIROS

	//INÍCIO NÚMEROS REAIS

	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 12300000.45
	fmt.Println(numeroReal2)

	numeroReal3 := 12234.67
	fmt.Println(numeroReal3)

	//FIM NÚMEROS REAIS

	//STRINGS

	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	//FIM STRINGS

	texto := 5
	fmt.Println(texto)

	var booleano1 bool = true
	fmt.Println(booleano1)

	var erro error = errors.New("Erro interno") // ERRORS é o nome do pacote
	fmt.Println(erro)

}
