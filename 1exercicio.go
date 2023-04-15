package main
//Crie um Array de inteiros com 3 elementos e imprima a soma dos valores armazenados no Array.
import (
	"fmt"
)

func main() {
	array := [3] int {1,2,3}
  soma := 0
  for _, valor := range array {
    soma +=valor 
  }
  fmt.Println("a soma do array é:", soma)
}
