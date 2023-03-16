package main

import "fmt"

func main() {
	var numero1 int
	var numero2 int
	var numero3 int
	var resultado int

	fmt.Scan(&numero1)
	fmt.Scan(&numero2)
	fmt.Scan(&numero3)
	resultado = numero1 + numero2 + numero3

	fmt.Println(resultado)

}
