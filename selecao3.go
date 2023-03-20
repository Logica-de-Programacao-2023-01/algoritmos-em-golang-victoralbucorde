package main

import "fmt"

func main() {
	var numero int
	fmt.Print("Escreva um número: ")
	fmt.Scan(&numero)

	if numero%2 == 0 {
		fmt.Println("O número é par")
	} else {
		fmt.Println("O número é impar")
	}
}
