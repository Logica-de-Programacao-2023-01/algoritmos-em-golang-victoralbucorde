package main

import "fmt"

func main() {
	var numero1 int
	var numero2 int
	var numero3 int
	fmt.Print("Escreva três números: ")
	fmt.Scan(&numero1)
	fmt.Scan(&numero2)
	fmt.Scan(&numero3)

	if numero1 < numero2 && numero1 < numero3 {
		fmt.Println("O menor número é", numero1)
	} else if numero2 < numero1 && numero2 < numero3 {
		fmt.Println("O menor número é", numero2)
	} else if numero3 < numero1 && numero3 < numero2 {
		fmt.Println("O menor número é", numero3)
	} else {
		fmt.Println("Os números são iguais")
	}
}
