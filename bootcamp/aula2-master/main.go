package main

import (
	"fmt"
)
var palavra = "hello"
func main(){
	fmt.Println(len(palavra))
	for pos, char := range palavra {
		fmt.Printf("character %c starts at byte position %d\n", char, pos)
	}

	

	
}

