package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero { // switch {} (outra forma de sintaxe)
	case 1: // case numero == 1: (outra forma de sintaxe)
		return "Domingo"
	case 2:
		return "Segunda-feira"
	case 3:
		return "Terça-feira"
	case 4:
		return "Quarta-feira"
	case 5:
		return "Quinta-feira"
	case 6:
		return "Sexta-feira"
	case 7:
		return "Sábado"
	default:
		return "Número Inválido"
	}
}

func main() {
	fmt.Println("Switch")

	dia := diaDaSemana(7)
	fmt.Println(dia)
}
