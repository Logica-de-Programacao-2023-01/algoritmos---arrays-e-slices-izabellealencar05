package main
//Crie um Slice de inteiros com o tamanho 5. Em seguida, solicite ao usuário que informe um número inteiro. Adicione esse número ao Slice apenas se ele não estiver presente. Imprima o Slice resultante.
import (
	"fmt"
)

func main() {
  slice := []int{1,2,3,4,5}
  var numero int

  fmt.Println("informe um numero inteiro:")
  fmt.Scan(&numero)

  estaNoSlice := false
    for _, valor := range slice{
      if valor == numero{
        estaNoSlice = true
        break
      }
    }
  if estaNoSlice == false{
    slice = append(slice, numero)
  }
  fmt.Print(slice)
	
}
