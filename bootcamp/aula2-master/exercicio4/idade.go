package main

import "fmt"

func main(){

	var employees = map[string]int{"Benjamin": 20, "Manuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 30}
	contador := 0
	fmt.Println(employees["Benjamin"])
	for key, element := range employees{
		if element > 21{
			fmt.Println(key)
			contador++
			
		}
	}
	fmt.Println(contador)

	employees["Frederico"] = 25
	delete(employees, "Pedro")
	fmt.Println(employees)




	
}