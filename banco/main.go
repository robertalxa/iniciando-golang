package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	conta1 := ContaCorrente{
		titular:       "Robert",
		numeroAgencia: 0001,
		numeroConta:   2000,
		saldo:         100,
	}
	conta2 := ContaCorrente{"Ariana Grande", 222, 333, 100}
	fmt.Println(conta1)
	fmt.Println(conta2)
}
