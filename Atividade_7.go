package main

import "fmt"

// Crie uma função que receba um slice de inteiros
//e retorne a média aritmética dos valores.

func main() {

	var z int

	x := []int{}

	fmt.Print("Digite números para serem inseridos no slice: ")
	fmt.Scan(&z)

	x = append(x, z)
	resultado := media(x)
	fmt.Println(resultado)

}
func media(x []int) int {

	resultado := 0

	for _, x := range x {
		resultado += x
	}
	return len(x) / resultado

}
