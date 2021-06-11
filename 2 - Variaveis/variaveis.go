package main

import "fmt"

func main() {
	var variavel1 string = "variável 1"
	variavel2 := "Variável 2"

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "lalal"
		variavel4 string = "lalalala"
	)
	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "Variável 5", "Variável 6"
	fmt.Println(variavel5, variavel6)

	const constante1 string = "constante 1"
	fmt.Println(constante1)

	//trocar valores:
	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)

}
