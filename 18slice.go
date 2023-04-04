package main

import "fmt"

//Escreva um programa que leia um número inteiro positivo n e gere um array com os n primeiros números primos.
func main() {
	var n int
	fmt.Print("digite o valo0e de n:")
	fmt.Scan(&n)

	var primo []int

	for i := 2; i <= n; i++ {
		primo := true
		for x := 2; x < i; x++ {
			if i%x == 0 {
				primo = false
				break
			}
		}
		if primo {
			primos = append(primos, i)
		}
	}
}
