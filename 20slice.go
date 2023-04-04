package main
//Crie um programa que leia um array de inteiros e verifique se ele está ordenado em ordem crescente
func main() {
	slice:=[]int{1,2,3,4,5,6,7,8}

	crescente := true
	for i:= 1, i< len (slice); i++{
		atual := slice[i]
		anterior := slice[i-1]

		if anterior>= atual{
			crescente = false
			break
		}
	}

}
