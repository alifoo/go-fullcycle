package main

import "fmt"

type Client struct {
	name string
}

func (c *Client) walked() {
	c.name = "alizinho"
	fmt.Printf("The client %s walked\n", c.name)
}

type Conta struct {
	saldo int
}

func NewConta() *Conta {
	return &Conta{saldo: 0}
}

func main() {
	ali := Client{
		name: "Alisson",
	}
	ali.walked()
	fmt.Println(ali.name)
}