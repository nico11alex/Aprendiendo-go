package main

import "fmt"

func generarNumeros(n int, c chan int) {
	for i := 1; i <= n; i++ {
		c <- i
	}
	close(c)
}

func main() {
	c := make(chan int)
	go generarNumeros(5, c)


	for i := range c{
		fmt.Println(i)
	}
	fmt.Println("Terminado.")
}
