package main

//Crie um Array bidimensional de inteiros com 3 linhas e 2 colunas. Solicite ao usuário que informe os valores de cada elemento da matriz. Em seguida, imprima a matriz resultante.
import "fmt"

func main() {
	var matriz [3][2]int

	for linha := 0; linha < len(matriz); linha++ {
		for coluna := 0; coluna < len(matriz[linha]); coluna++ {
			var x int
			fmt.Println("digite o numero da linha", linha, "e da coluna", coluna)
			fmt.Scan(&x)
			matriz[linha][coluna] = x
		}
		fmt.Println(matriz)
	}

}