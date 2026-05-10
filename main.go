package main

import (
	"errors"
	"fmt"
	"sync"
)

type ErrorPermiso struct {
	Nombre string
}

type Resultado struct {
	Nombre string
	Err    error
}

var ErrorArchivoNoEncontrado = errors.New("Archivo no encontrado:")

func verificarArchivo(wg *sync.WaitGroup,nombre string, res chan<- Resultado) {
	defer wg.Done()
	if nombre != "main.go" {
		res <- Resultado{Err: fmt.Errorf("%w %s", ErrorArchivoNoEncontrado, nombre)}
		return
	}
	res <- Resultado{Nombre: nombre}
}

func main() {
	var wg sync.WaitGroup
	c := make(chan Resultado, 4)
	archivos := []string{"main.go", "slice.py", "gotes.js", "libro.php"}

	
	for _, archivo := range archivos {
		wg.Add(1)
		go verificarArchivo(&wg,archivo, c)
	}

	wg.Wait()
	for i := 0; i < 4; i++ {
		r := <-c

		if errors.Is(r.Err, ErrorArchivoNoEncontrado) {
			fmt.Println(r.Err)
		} else {
			fmt.Println("Archivo encontrado:", r.Nombre)
		}

	}

}
