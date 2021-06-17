package main

import "fmt"

func main() {
	fmt.Println("maps")

	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}
	fmt.Println(usuario)

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "João",
			"último":   "Pedro",
		},
		"curso": {
			"nome":   "Engenharia",
			"campus": "Campus 1",
		},
	}
	fmt.Println(usuario2)
	delete(usuario2, "nome") //DELETE
	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string{
		"nome": "´Gémeos",
	}
	fmt.Println(usuario2)

}
