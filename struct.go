package main

import (
	"fmt"
	"math"
)

// Crie uma struct chamada Circulo com o campo "raio".
// Escreva uma função que receba um Circulo como parâmetro e calcule a área do círculo
// (área = pi * raio * raio). Dica: use a constante math.Pi para representar o número pi.
func main() {
	var x float64
	fmt.Print("Digite o valor do seu raio:")
	fmt.Scan(&x)
	r := circulo{raio: x}
	fmt.Printf("A area do circulo é: %.2f", area(r.raio))
}

type circulo struct {
	raio float64
}

func area(raio float64) float64 {
	area := math.Pi * math.Pow(raio, 2)
	return area
}
