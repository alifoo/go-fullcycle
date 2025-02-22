package main

import "fmt"

type Client struct {
	Nome string
	Idade int
	Ativo bool
}

func main() {
	ali := Client{
		Nome: "Alisson",
		Idade: 22,
		Ativo: true,
	}

	fmt.Println(ali)
}