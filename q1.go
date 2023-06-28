package main

import (
	"fmt"
	"math"
)

//Função que recebe um struct de um Círculo e calcula sua área.

type Circulo struct {
	raio float64
}

func calcularArea(c Circulo) float64 {
	return math.Pi * c.raio * c.raio
}

func main() {
	circulo := Circulo{raio: 5.0}
	area := calcularArea(circulo)
	fmt.Printf("Área do círculo: %.2f\n", area)
}
