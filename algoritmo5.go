package main

import "fmt"

func main() {
	var idade int
	var idadeDias int

	fmt.Scan(&idade)
	idadeDias = idade * 365
	fmt.Println("Idade em dias:", idadeDias)
}
