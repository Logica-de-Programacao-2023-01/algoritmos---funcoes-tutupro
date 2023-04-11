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

package main

import (
	"fmt"
)

func main() {
	var x string

	str := []string{}

	fmt.Print("Digite uma palavra aleatória: ")
	fmt.Scan(&x)

	u := append(str, x)

	resultado := countVowels(str, str)

	fmt.Println(resultado, u)

}

func countVowels(str string) int {
	count := 0
	for _, char := range str {
		switch char {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			count++
		}
	}
	return count
}

Atividade 3º :

