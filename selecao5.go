package main

import "fmt"

func main() {
	var numero int
	fmt.Print("Escreva um número: ")
	fmt.Scan(&numero)

	if numero%3 == 0 && numero%5 == 0 {
		fmt.Println("O número é multiplo de 3 e 5")
	} else {
		fmt.Println("O número não é multiplo de 3 e 5")
	}
}
