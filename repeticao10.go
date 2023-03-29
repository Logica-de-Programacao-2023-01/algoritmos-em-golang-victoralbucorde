package main

import "fmt"

func main() {

	var maior, numero, contador float64

	for {
		fmt.Scan(&numero)
		if numero == 0 {
			break
		}
		if numero >= maior {
			maior = numero
		}
		contador++
	}
	if contador > 0 {
		fmt.Println("O maior dos", contador, "numeros é ", maior)
	} else {
		fmt.Println("Nenhum numero foi digitado")
	}

}
