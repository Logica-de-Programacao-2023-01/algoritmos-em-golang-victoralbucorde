package main

import "fmt"

func main() {
	var salario float64

	fmt.Scan(&salario)

	fmt.Println("Novo salario:", salario*1.15)
}
