package main

func main() {
	a := 10
	var ponteiro *int = &a
	*ponteiro = 20
	b := &a

	*b = 30
	println(a)

}
