package main

import (
	"fmt"
	"sync"
	"time"
)

func trabajador(id int, wg *sync.WaitGroup, resultados chan<- string) {
	defer wg.Done()

	start := time.Now()
	fmt.Printf("Trabajador %d empezó\n", id)

	time.Sleep(time.Duration(id) * 500 * time.Millisecond)

	duracion := time.Since(start)

	mensaje := fmt.Sprintf("Trabajador %d terminó en %v", id, duracion)
	resultados <- mensaje
}

func main() {
	var wg sync.WaitGroup
	resultados := make(chan string, 5)

	fmt.Println("Iniciando trabajadores...")

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go trabajador(i, &wg, resultados)
	}

	wg.Wait()
	close(resultados)

	fmt.Println("\n=== Resultados ===")
	for msg := range resultados {
		fmt.Println("→", msg)
	}

	fmt.Println("\n¡Todo terminado!")
}
