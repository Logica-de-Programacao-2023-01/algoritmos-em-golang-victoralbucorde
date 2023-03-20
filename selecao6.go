package main

import "fmt"

func main() {
	var numero1 int
	var numero2 int
	fmt.Print("Escreva dois números: ")
	fmt.Scan(&numero1)
	fmt.Scan(&numero2)

	if numero1 < 0 || numero2 < 0 {
		fmt.Println("A soma dos dois números é", numero1+numero2)
	} else {
		fmt.Println("O produto dos dois números é", numero1*numero2)
	}
}
