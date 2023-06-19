package main

import "fmt"

// retorna mais de um valor
func sum(x int, y int) (int, bool) {
	return x + y, true
}

func main() {
	// O : declara a variavel, usa somente na 1a vez
	name := "Antonio" 
	name = "Pedro"
	fmt.Println(name, "outro texto")
	println(sum(2, 5))

	newSum()
}

func newSum() {
	result, status := sum(1, 5)
	println(result, status)
}