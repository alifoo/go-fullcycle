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

type Pessoa interface {
	Deactivate()
}

func Deactivation(pessoa Pessoa) {
	pessoa.Deactivate()
}

func main() {
	ali := Client{
		Nome: "Alisson",
		Idade: 22,
		Ativo: true,
	}

	ali.Address.Cidade = "Curitiba"
	Deactivation(&ali)

	fmt.Println(ali)
}