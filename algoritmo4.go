package main

import "fmt"

func main() {
	var numero1 float64
	var numero2 float64
	var numero3 float64
	var media float64

	fmt.Scan(&numero1)
	fmt.Scan(&numero2)
	fmt.Scan(&numero3)
	media = (numero1*2 + numero2*3 + numero3*5) / 10

	fmt.Println(media)

}
