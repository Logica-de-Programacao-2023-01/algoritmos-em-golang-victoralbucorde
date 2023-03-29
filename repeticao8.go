package main

import "fmt"

func main() {

	var numero int
	fmt.Print("Escreva um numero")
	fmt.Scan(&numero)

	for i := 1; i <= numero; i++ {
		if numero%i == 0 {
			fmt.Println(i)
		}
	}
}
