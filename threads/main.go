package main

import (
	"fmt"
	"time"
)

func contador(tipo string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%s: %d\n", tipo, i)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	// runThreads()
	// canalThread()
	// testThread()

	msg := make(chan int)

	for i := 0; i < 50; i++ {
		go workerCalculadora(i, msg)
	}
	// go workerCalculadora(1, msg)
	// go workerCalculadora(2, msg)
	// go workerCalculadora(3, msg)

	for i := 0; i < 100; i++ {
		msg <- i
	}
}

func runThreads() {
	go contador("a")
	go contador("b")
	time.Sleep(time.Second * 10)
}

func canalThread() {
	canal := make(chan string)

	go processa(canal)

	result := <-canal
	fmt.Println(result)
}

func processa(canal chan string) {
	canal <- "processado"
}

func testThread() {
	queue := make(chan int)
	go func() {
		i := 0
		for {
			queue <- i // Envia o valor para o canal (canal cheio)
			i++
			time.Sleep(time.Second)
		}
	}()

	for x := range queue {
		fmt.Println(x) // Recebe o valor do canal (canal vazio)
	}
}

func workerCalculadora(workerId int, msg chan int) {
	for res := range msg {
		fmt.Printf("Worker %d recebeu %d\n", workerId, res)
		time.Sleep(time.Second)
	}
}


