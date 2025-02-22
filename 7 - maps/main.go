package main

import "fmt"

func main() {
	salarios := map[string]int{"Alisson": 5500, "Myllena": 7500}
	fmt.Println(salarios["Alisson"])

	for nome, salario := range salarios {
		fmt.Printf("The salary of %s is %d\n", nome, salario)
	}
}