package main

import "fmt"

// Crie uma struct ramda Pessoa com os campos "nome", "idade" e "endereço". O
// campo "endereço" deve ser uma outra struct com os campos "rua", "número",
// "cidade" e "estado". Escreva uma função que receba uma Pessoa como parâmetro e
// imprima seu endereço completo.
func main() {
	p := pessoa{
		nome:  "kaynan",
		idade: 21,
		endereco: endereco{
			rua:    "shin",
			numero: 209,
			cidade: "brasilia",
			estado: "df",
		},
	}
	fmt.Print(p)
}

type endereco struct {
	rua    string
	numero int
	cidade string
	estado string
}
type pessoa struct {
	nome  string
	idade int
	endereco
}
