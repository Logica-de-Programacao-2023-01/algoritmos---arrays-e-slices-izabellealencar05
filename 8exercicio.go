package main
//Crie um Slice de strings com tamanho 8 e solicite ao usuário que informe um valor. Remova todas as ocorrências desse valor no Slice e imprima o resultado.
import (
	"fmt"
)

func main() {
	slice := [] string {"por", "que", "boz", "faz", "dez", "liz", "lin", "pra"}
  var valor string
  
  fmt.Println("diigte algo a ser removido do slice", slice)
  fmt.Scan(&valor)

  for i := 0; i<len(slice); i++{
    if slice[i] == valor{
      slice = append(slice[:i], slice[i+1:]...)
      i--
    }
  } 
  fmt.Print("slice resultante:", slice)
}
