package main
//Crie um Array de floats com 6 elementos. Solicite ao usuário que informe um número que será adicionado em todas as posições do Array. Imprima o Array resultante.
import (
	"fmt"
)

func main() {
array := [6] float64 {1.2, 2.5, 3.6, 4.5, 5.6, 6.7}
var valor float64

  fmt.Println("informe um valor para ser adicionado no array")
  fmt.Scan(&valor)

  for i := 0; i < len(array); i++ {
    array[i] += valor
  }
  fmt.Println("o array com o numero adicionado é: ", array)
}
