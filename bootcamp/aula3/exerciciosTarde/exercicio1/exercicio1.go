package main

import "fmt"

type Estudante struct{
	Nome	     string
	Sobrenome    string
	Rg		     string
	DataAdmissao string

}

func detalhes(estudante Estudante){
	fmt.Println("Nome:" ,estudante.Nome, "Sobrenome:", estudante.Sobrenome, "Rg:", estudante.Rg, "Data de Admissão: ", estudante.DataAdmissao )
}
func main(){
	estudante1 := Estudante {
	Nome:	"Vivian",
	Sobrenome: "Sanches",
	Rg: "20421878",
	DataAdmissao: "16052022",
	}

	detalhes(estudante1)
}