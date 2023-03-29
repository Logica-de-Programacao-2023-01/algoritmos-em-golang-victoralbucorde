package main

import "fmt"

func main() {

	var numero, media, divisor, soma float64

	fmt.Println("Escreva varios numeros para fazer uma media aritmetica entre eles quando terminar digite 0")
	for {
		fmt.Scan(&numero)
		if numero == 0 {
			break
		}

		divisor++
		soma += numero
	}
	if divisor > 0 {
		media = soma / divisor
		fmt.Printf("Sua media é %.2f\n", media)
	} else {
		fmt.Println("Nenhum numero foi digitado")
	}

}
