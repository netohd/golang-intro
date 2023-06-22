package main

import (
	"fmt"
	"time"
)

func worker(workerId int, data chan int) {
	for x := range data {
		// Quando o x é resolvido o canal é limpado
		fmt.Printf("Worker %d received %d\n", workerId, x)
		time.Sleep(time.Second)
	}
}

func main() {
	// Canal usado p/ threads se comunicarem
	channel := make(chan int)
	
	// Cria as outras threads
	for i:= 0; i < 1000; i++ {
		go worker(i, channel)
	}

	// Cria as requisições a serem resolvidas no canal
	for i:= 0; i < 100000; i++ {
		channel <- i
	}
}