package main

func sum(x int, y int) (int, bool) {
	return x + y, true
}

func main() {
	// O : declara a variavel, usa somente na 1a vez
	name := "Antonio" 
	name = "Pedro"
	println(name)
	println(sum(2, 5))

	newSum()
}

func newSum() {
	result, status := sum(1, 5)
	println(result, status)
}