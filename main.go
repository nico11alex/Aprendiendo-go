package main

import (
	"fmt"
	"sync"
	"time"
)

// FAN-IN

func productor(id int, frases []string, c chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	var frase string
	for _, f := range frases {
		frase = fmt.Sprintf("Producto %d: %s", id, f)
		c <- frase
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {

	c := make(chan string)

	var wg sync.WaitGroup

	cadena1 := []string{"Hola", "Mundo", "Go"}
	cadena2 := []string{"Concurrencia", "es", "poderosa"}

	wg.Add(1)
	go productor(1, cadena1, c, &wg)
	wg.Add(1)
	go productor(2, cadena2, c, &wg)

	go func() {
		wg.Wait()
		close(c)
	}()

	for msg := range c {
		fmt.Println(msg)
	}

}
