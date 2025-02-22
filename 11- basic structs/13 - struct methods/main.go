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

func (c *Client) Deactivate() {
	c.Ativo = false
}

func main() {
	ali := Client{
		Nome: "Alisson",
		Idade: 22,
		Ativo: true,
	}

	ali.Address.Cidade = "Curitiba"
	ali.Deactivate()

	fmt.Println(ali)
}