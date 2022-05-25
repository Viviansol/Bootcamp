package main

import (
	"fmt"
)
var palavra = "hello"
func letras(){
	fmt.Println(len(palavra))
	for pos, char := range palavra {
		fmt.Printf("character %c starts at byte position %d\n", char, pos)
	}	
}