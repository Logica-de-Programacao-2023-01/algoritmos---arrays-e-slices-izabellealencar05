package main
//Crie um Array de inteiros com 10 elementos. Em seguida, solicite ao usuário que informe um valor e verifique se esse valor está presente no Array. Imprima o resultado da busca.
import (
	"fmt"
)

func main() {
  var numero int
	array:= [10] int {1,2,3,4,5,6,7,8,9,10}

  fmt.Println("informe um numero:")
  fmt.Scan(&numero)

  encontrado := false
  for _, elemento := range array{
    if elemento == numero{
      encontrado = true
      break
    }
  }
  if encontrado{
    fmt.Print("o valor esta presente no array")
  } else {
    fmt.Print("o valor nao esta presente no array")
  }
}
