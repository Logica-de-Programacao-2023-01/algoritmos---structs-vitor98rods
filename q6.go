package main

import (
	"fmt"
	"time"
)

//Função que recebe um struct de funcionário que reajusta o salário do próprio a partir de funções que calculam em a porcentegem e tempo serviço do funcionário.

type Funcionario struct {
	Nome    string
	Salario float64
	Idade   int
}

func (f *Funcionario) AumentarSalario(porcentagem float64) {
	f.Salario *= (1 + porcentagem/100)
}

func (f *Funcionario) DiminuirSalario(porcentagem float64) {
	f.Salario *= (1 - porcentagem/100)
}

func (f *Funcionario) TempoServico() int {
	idadeInicioTrabalho := 18
	idadeAtual := time.Now().Year() - f.Idade
	tempoServico := idadeAtual - idadeInicioTrabalho
	return tempoServico
}

func main() {
	funcionario := Funcionario{
		Nome:    "Jonas",
		Salario: 3000,
		Idade:   25,
	}

	fmt.Println("Salário atual:", funcionario.Salario)
	funcionario.AumentarSalario(10)
	fmt.Println("Salário após aumento:", funcionario.Salario)
	funcionario.DiminuirSalario(5)
	fmt.Println("Salário após redução:", funcionario.Salario)

	tempoServico := funcionario.TempoServico()
	fmt.Println("Tempo de serviço:", tempoServico, "anos")
}
