package main

import "os"

func main(){
	data, err := os.ReadFile("./products.txt")

	
	println(data)
	print(err)
}