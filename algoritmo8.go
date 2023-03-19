package main

import "fmt"

func main() {
	var dias int
	var diaria float64
	var salario float64

	fmt.Scan(&dias)
	fmt.Scan(&diaria)
	salario = float64(dias) * diaria

	fmt.Println(salario)

}
