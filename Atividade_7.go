Codigo De funções

Atividade 1º:

package main

import "fmt"

func main() {
	var z int
	slice := []int{}
	fmt.Print("Digite números para serem inseridos no slice: ")
	fmt.Scan(&z)

	x := append(slice, z)
	println(x)
	media := mediaAritmetica(slice)
	fmt.Println(media)
}

func mediaAritmetica(slice []int) float64 {
	soma := 0
	for _, valor := range slice {
		soma += valor
	}
	return float64(soma) / float64(len(slice))
}

Atividade 2º:

