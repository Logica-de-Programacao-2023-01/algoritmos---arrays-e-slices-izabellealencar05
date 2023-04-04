package main

import "fmt"

//Crie um Slice de inteiros com tamanho 8 e solicite ao usuário que informe dois índices de elementos que deverão ser trocados de posição. Imprima o Slice resultante.

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	var x, y int
	fmt.Print("digite o primeiro elemento:")
	fmt.Scan(&x)
	fmt.Print("digite o segundo elemento:")
	fmt.Scan(&y)

	aux := slice[x]
	slice[x] = slice[y]
	slice[y] = aux

}
