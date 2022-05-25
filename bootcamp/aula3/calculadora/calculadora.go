package main 

import "fmt"

const(
	Soma = "+"
	Sub = "-"
	Mut = "*"
	Div = "/"
)

func soma(values ...float64) float64{
	var resultado float64
	for _, value := range values{
		resultado += value
	}
	return resultado
}

func operador(valor1, valor2 float64, operador string) float64{
	switch operador{
	case Soma:
		return valor1 + valor2
	case Sub:
		return valor1 - valor2
	case Mut:
		return valor1 * valor2
	case Div:
		if valor2 != 0{
			return valor1/valor2
		}
	}
	return 0
}


func main(){
	fmt.Println(operador(2,2,Soma))
	fmt.Println(soma(10,20,85,93,1000))
}
