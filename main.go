package main

import (
	"fmt"
	"net/http"
)

type Order struct {
	id         int
	price      float32
	address    string
	clientName string
}

func main() {
	// Inicia servidor e chama função para rota
	http.HandleFunc("/", home)
	http.HandleFunc("/order", order)
	http.ListenAndServe(":8000", nil)
}

func order(w http.ResponseWriter, r *http.Request) {
	testOrder := Order{
		id:         1,
		price:      29.90,
		address:    "Rua do teste",
		clientName: "Joao",
	}

	w.Write([]byte(testOrder.address))
}

func home(w http.ResponseWriter, r *http.Request) {
	// O : declara a variavel, usa somente na 1a vez
	name := "Antonio"
	name = "Pedro"
	fmt.Println(name, "second arg")
	println(sum(2, 5))
	w.Write([]byte("name"))
}

// retorna mais de um valor
func sum(x int, y int) (int, bool) {
	return x + y, true
}

func newSum() {
	result, status := sum(1, 5)
	println(result, status)
}
