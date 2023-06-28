package main

import "fmt"

//Função que recebe um struct "filme" que modifica suas avaliações a adciona outras.

type Filme struct {
	Titulo     string
	Diretor    string
	Ano        int
	Avaliacoes []int
}

func (f *Filme) AdicionarAvaliacao(avaliacao int) {
	f.Avaliacoes = append(f.Avaliacoes, avaliacao)
}

func (f *Filme) RemoverAvaliacao(indice int) {
	if indice >= 0 && indice < len(f.Avaliacoes) {
		f.Avaliacoes = append(f.Avaliacoes[:indice], f.Avaliacoes[indice+1:]...)
	}
}

func (f *Filme) CalcularMediaAvaliacoes() float64 {
	total := 0

	for _, avaliacao := range f.Avaliacoes {
		total += avaliacao
	}

	media := float64(total) / float64(len(f.Avaliacoes))
	return media
}

func (f *Filme) ImprimirInformacoes() {
	fmt.Printf("Título: %s\n", f.Titulo)
	fmt.Printf("Diretor: %s\n", f.Diretor)
	fmt.Printf("Ano: %d\n", f.Ano)
	fmt.Printf("Média de Avaliações: %.2f\n", f.CalcularMediaAvaliacoes())
}

func main() {
	filme := Filme{
		Titulo:     "Pulp Fiction",
		Diretor:    "Quentin Tarantino",
		Ano:        1994,
		Avaliacoes: []int{5, 4, 4, 5, 3},
	}

	filme.ImprimirInformacoes()

	filme.AdicionarAvaliacao(4)
	filme.RemoverAvaliacao(2)

	filme.ImprimirInformacoes()
}
