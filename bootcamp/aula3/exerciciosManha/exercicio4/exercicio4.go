package main
import (
	"fmt"
)


func opMinimo(notas... float64) float64{
	var min float64 = 99999
	for _, e := range notas {
		if e < min {
			min = e
		}
	}
	return min
}

func opMedia (notas... float64) float64{
	var resultado float64
	
	for _, value := range notas{
		resultado += value
	}
	var qnt int = len(notas)
	return resultado/float64(qnt)
}

func opMaximo(notas... float64)float64{
	var max float64 = -1
	for _, e := range notas {
		if e > max {
			max = e
		}
	}
	return max
}

func alunodados(dado string ) func(notas... float64) float64{
	switch(dado){
	case "maximo":
		return opMaximo
	case "minimo":
		return opMinimo
	case "media":
		return opMedia
	}
	return nil
	
}


func main(){
oper := alunodados("media")
r:= oper(4,4,45,5)
fmt.Println(r)
}