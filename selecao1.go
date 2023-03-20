package main

import "fmt"

func main() {
	var numero1 int
	var numero2 int
	fmt.Print("Escreva dois numeros: ")
	fmt.Scan(&numero1)
	fmt.Scan(&numero2)

	if numero1 > numero2 {
		fmt.Println("O maior número é", numero1)
	} else if numero1 < numero2 {
		fmt.Println("O maior número é", numero2)
	} else {
		fmt.Println("Os números são iguais")
	}
}
