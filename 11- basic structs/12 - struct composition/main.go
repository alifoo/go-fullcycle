package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero int
	Cidade string
	Estado string
}

type Client struct {
	Nome string
	Idade int
	Ativo bool
	Address Endereco
}

func main() {
	ali := Client{
		Nome: "Alisson",
		Idade: 22,
		Ativo: true,
	}

	ali.Address.Cidade = "Curitiba"

	fmt.Println(ali)
}