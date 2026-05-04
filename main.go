package main

import (
	"fmt"
	"sync"
	"time"
)

func trabajador(id int, tareas <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range tareas {
		fmt.Printf("Trabajador %d procesando: %s\n", id, i)
		time.Sleep(500 * time.Millisecond)
	}

}

func main() {
	c := make(chan string, 3)
	var wg sync.WaitGroup

	wg.Add(1)
	go trabajador(1, c, &wg)
	wg.Add(1)
	go trabajador(2, c, &wg)
	wg.Add(1)
	go trabajador(3, c, &wg)
	tarea := []string{"A", "B", "C", "D", "E"}

	for _, s := range tarea {
		c <- s
	}

	close(c)
	wg.Wait()
	fmt.Println("Todas las tareas completadas.")
}
