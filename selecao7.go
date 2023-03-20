package main

import "fmt"

func main() {
	var salario float64
	fmt.Print("Escreva o salário em R$")
	fmt.Scan(&salario)

	if salario >= 1000 {
		fmt.Printf("O salário com aumento é igual a R$%.2f", salario*1.05)
	} else {
		fmt.Printf("O salário com aumento é igual a R$%.2f", salario*1.1)
	}
}
