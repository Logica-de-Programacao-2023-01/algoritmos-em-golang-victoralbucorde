package main

import "fmt"

func main() {
	var peso float64
	var altura float64
	fmt.Print("Digite o peso e altura de uma pessoa respectivamente: ")
	fmt.Scan(&peso)
	fmt.Scan(&altura)
	imc := peso / (altura * altura)

	if imc < 18.5 {
		fmt.Println("Abaixo do peso")
	} else if imc > 24.9 {
		fmt.Println("Acima do peso")
	} else {
		fmt.Println("Peso ideal")
	}
}
