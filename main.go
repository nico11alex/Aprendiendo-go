package main

import (
	"fmt"
	"sync"
)

func echo(c chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for s := range c {
		fmt.Printf("Eco: %s\n", s)
	}
}

func main() {
	c := make(chan string)
	var wg sync.WaitGroup

	wg.Add(2)
	go echo(c, &wg)
	go echo(c, &wg)
	enviar := []string{"1", "2"}
	for i := 0; i <= 10; i++ {
		for _, s := range enviar {
			c <-s
		}
	}
	close(c)
	wg.Wait()
	fmt.Println("Terminado.")
}
