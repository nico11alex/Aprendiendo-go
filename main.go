package main

import (
	"fmt"
	"sync"
)

func echo(c chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for s := range c{
		fmt.Printf("Eco: %s\n",s)
	}
}

func main() {
	c := make(chan string)
	var wg sync.WaitGroup

	wg.Add(1)
	go echo(c, &wg)

	enviar := []string{"Hola", "Mundo", "Go"}
	for _,s := range enviar{
		c <- s	
	}
	
	close(c)
	wg.Wait()
	fmt.Println("Terminado.")
}
