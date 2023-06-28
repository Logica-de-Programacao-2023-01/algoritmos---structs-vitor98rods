package main

import "fmt"

//Função que Modifica sons de struct de animais.

type Animal struct {
	Nome    string
	Especie string
	Idade   int
	Som     string
}

func (a *Animal) AlterarSom(novoSom string) {
	a.Som = novoSom
}

func (a Animal) ImprimirInformacoes() {
	fmt.Printf("Nome: %s\n", a.Nome)
	fmt.Printf("Espécie: %s\n", a.Especie)
	fmt.Printf("Idade: %d\n", a.Idade)
	fmt.Printf("Som: %s\n", a.Som)
}

func main() {
	animal := Animal{
		Nome:    "Cachorro",
		Especie: "Dobermann",
		Idade:   5,
		Som:     "Latido",
	}

	animal.ImprimirInformacoes()

	animal.AlterarSom("Au au")

	animal.ImprimirInformacoes()
}
