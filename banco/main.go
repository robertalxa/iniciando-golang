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

	conta2 := ContaCorrente{
		titular:       "Robert",
		numeroAgencia: 0001,
		numeroConta:   2000,
		saldo:         100,
	}
	fmt.Println(conta1 == conta2) // O Go entende que tem que comparar o conteúdo e que o mesmo é de fato igual

	// fmt.Println(conta1)
	// fmt.Println(conta2)

	// conta2 := ContaCorrente{"Ariana Grande", 222, 333, 100}
	// fmt.Println(conta2)

	// var contaBieber *ContaCorrente

	// contaBieber = new(ContaCorrente)
	// contaBieber.titular = "Bieber"
	// contaBieber.saldo = 5000000000
	// fmt.Println(*contaBieber)
}
