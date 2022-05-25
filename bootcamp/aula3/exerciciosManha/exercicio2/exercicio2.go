package main
import "fmt"

func media(notas...int)int{
	var resultado int
	
	for _, value := range notas{
		resultado += value
	}
	var qnt int = len(notas)
	return resultado/qnt
}


func main(){
	fmt.Println(media(10,14,98,64))
}