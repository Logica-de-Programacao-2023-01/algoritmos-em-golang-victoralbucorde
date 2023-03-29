package main

import "fmt"

func main() {

	var numero float64
	fmt.Print("Escreva um numero: ")
	fmt.Scan(&numero)

	for i := 1.0; i <= 10; i++ {
		fmt.Println(i * numero)
	}
}
