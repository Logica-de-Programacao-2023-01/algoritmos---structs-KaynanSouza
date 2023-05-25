package main

import "fmt"

// Crie uma struct chamada Triângulo com os campos "base" e "altura". Escreva uma
// função que receba um Triângulo como parâmetro e calcule a área do triângulo
// (área = base * altura / 2).
func main() {
	var altura, base float64
	fmt.Print("Digite a altura do triangulo:")
	fmt.Scan(&altura)
	fmt.Print("Digite a base do triangulo:")
	fmt.Scan(&base)
	t := triangulo{altura: altura, base: base}
	area := area(t.altura, t.base)
	fmt.Printf("A altura do triangulo é: %.2f", area)
}

type triangulo struct {
	base   float64
	altura float64
}

func area(base, altura float64) float64 {
	return (base * altura) / 2
}
