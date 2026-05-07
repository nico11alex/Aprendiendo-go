package main

import (
	"fmt"
	"time"
)

func buscarRespuesta(resultado chan<- string) {
	time.Sleep(3 * time.Second)
	resultado <- "Datos encontrados"
	close(resultado)
}

func main() {
	c := make(chan string)
	go buscarRespuesta(c)

	select{
	case msg := <-c :
		fmt.Printf("Éxito: %s\n",msg)
	case <-time.After(2*time.Second):
		fmt.Println("Tiempo agotado. La búsqueda tardó demasiado.")
	}
	fmt.Println("Programa finalizado.")
}
