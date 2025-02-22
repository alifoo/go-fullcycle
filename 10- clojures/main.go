package main

import "fmt"

func main() {
	total := func() int {
		return sum(1,2,3,325,2435,3456,543,6347,34,5785345,654784) * 2
	}()
	fmt.Println(total)
}


func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}