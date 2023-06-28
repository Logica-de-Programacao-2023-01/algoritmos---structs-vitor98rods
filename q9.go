package main

import "fmt"

//Função que recebe um struct "aluno" e retorna "notas" como um slice para representar cada matéria.

type Aluno struct {
	Nome  string
	Idade int
	Notas []float64
}

func (a *Aluno) AdicionarNota(nota float64) {
	a.Notas = append(a.Notas, nota)
}

func (a *Aluno) RemoverNota(indice int) {
	if indice >= 0 && indice < len(a.Notas) {
		a.Notas = append(a.Notas[:indice], a.Notas[indice+1:]...)
	}
}

func (a *Aluno) CalcularMedia() float64 {
	total := 0.0

	for _, nota := range a.Notas {
		total += nota
	}

	media := total / float64(len(a.Notas))
	return media
}

func (a *Aluno) ImprimirInformacoes() {
	fmt.Printf("Nome: %s\n", a.Nome)
	fmt.Printf("Idade: %d\n", a.Idade)
	fmt.Printf("Média: %.2f\n", a.CalcularMedia())
}

func main() {
	aluno := Aluno{
		Nome:  "João",
		Idade: 20,
		Notas: []float64{8.5, 7.9, 9.1},
	}

	aluno.ImprimirInformacoes()

	aluno.AdicionarNota(8.0)
	aluno.RemoverNota(1)

	aluno.ImprimirInformacoes()
}
