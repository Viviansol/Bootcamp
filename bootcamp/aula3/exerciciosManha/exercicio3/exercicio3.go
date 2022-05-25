package main 

import "fmt"


func salarioA(horas, salario float64)float64{
	if horas > 160{
		return salario * horas + (0.50*salario)* horas
	}else{
		return salario*horas
	}
}

func salarioB(horas, salario float64)float64{
	if horas >160{
		return salario * horas + (0.50*salario)* horas
	}else{
		return salario*horas
	}
}
func salarioC(horas, salario float64)float64{
		return salario*horas
}

func main(){
fmt.Println(salarioC(162, 1000))
fmt.Println(salarioB(176, 1500))
fmt.Println(salarioA(172, 3000))
}