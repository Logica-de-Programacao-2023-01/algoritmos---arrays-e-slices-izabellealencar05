//Crie um Array de floats com 4 elementos e calcule o produto dos valores armazenados no Array.
package main
import "fmt"
func main() {
  array := [4] float64 {1.5,2.5,3.5,4.5}

  produto := 1.0
  
  for _, valor := range array {
    produto*= valor
  }
  fmt.Println("o produto dos valores do array é:" produto)
}