package main


import(
	"os"
	"fmt"
)

type Produto struct{
	Nome	   string
	Preco      float64
	Id		   int
	Quantidade int
}

func detalhes(produto Produto) string {

	nome := produto.Nome
	preco := produto.Preco
	Id := produto.Id
    quant := produto.Quantidade

	dados := fmt.Sprint("\n Nome: ", nome, "\n Preco: ", preco, "\n ID ", Id, "\n Quant: ", quant)

	return dados
}


func write( produto Produto ){

	
	dado := detalhes(produto)
	product :=  []byte (dado)
	err := os.WriteFile("./products.txt", product, 0644)
	fmt.Println(err)
}

func main(){
	p1 := Produto{
		Nome: "Pincel",
		Preco: 10.00,
		Id: 1,
		Quantidade: 100,
	}
	p2 := Produto{
		Nome: "Lapis de cor",
		Preco: 1000.00,
		Id: 2,
		Quantidade: 100,
	}
	write(p1)
	write(p2)
	
}