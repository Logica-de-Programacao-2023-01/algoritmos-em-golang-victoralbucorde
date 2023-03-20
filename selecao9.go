package main

import "fmt"

func main() {
	var numero1 float64
	var numero2 float64
	var numero3 float64
	fmt.Print("Escreva três números: ")
	fmt.Scan(&numero1)
	fmt.Scan(&numero2)
	fmt.Scan(&numero3)

	if numero1 < numero2 && numero2 < numero3 {
		fmt.Println("Em ordem crescente: ", numero1, numero2, numero3)
	} else if numero1 < numero3 && numero3 < numero2 {
		fmt.Println("Em ordem crescente: ", numero1, numero3, numero2)
	} else if numero2 < numero1 && numero1 < numero3 {
		fmt.Println("Em ordem crescente: ", numero2, numero1, numero3)
	} else if numero2 < numero3 && numero3 < numero1 {
		fmt.Println("Em ordem crescente: ", numero2, numero3, numero1)
	} else if numero3 < numero1 && numero1 < numero2 {
		fmt.Println("Em ordem crescente: ", numero3, numero1, numero2)
	} else if numero3 < numero2 && numero2 < numero1 {
		fmt.Println("Em ordem crescente: ", numero3, numero2, numero1)
	} else {
		fmt.Println("Os números são iguais")
	}
}
