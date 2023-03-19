package main

import "fmt"

func main() {
	var peso float64
	var libra float64 = 2.204

	fmt.Println("Escreva o peso em quilos:")
	fmt.Scan(&peso)

	fmt.Printf("Peso em libras: %.2f\n", peso*libra)

}
