package main

import "fmt"

func main() {
	var altura float64
	var peso float64
	var imc float64
	fmt.Scan(&peso)
	fmt.Scan(&altura)

	imc = peso / (altura * altura)

	fmt.Printf("IMC = %.1f\n", imc)

}
