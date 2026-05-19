package main

import (
	"fmt"
	"sync"
	"time"
)

type TareaResultado struct {
	ID        int
	Tarea     string
	Resultado string
	Error     error
	Tiempo    time.Duration
}

func procesarTarea(id int, c chan<- TareaResultado){
	start := time.Now()
	tareaNombre := fmt.Sprintf("Tarea-%d", id)

	// Simular trabajo
	time.Sleep(time.Duration(id) * 300 * time.Millisecond)

	// Algunas tareas fallan
	if id == 2 || id == 4 {
		c <-TareaResultado{
			ID:     id,
			Tarea:  tareaNombre,
			Error:  fmt.Errorf("fallo en la conexión con el servidor"),
			Tiempo: time.Since(start),
		}
		return
	}

	c <-TareaResultado{
		ID:        id,
		Tarea:     tareaNombre,
		Resultado: "Procesado correctamente",
		Tiempo:    time.Since(start),
	}
}

func main() {
	var wg sync.WaitGroup
	c := make(chan TareaResultado)

	for i:=1;i <= 5;i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			procesarTarea(id,c)
		}(i)
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	for  res := range c{
		if res.Error != nil{
			fmt.Printf("Error en la %s: %v\n",res.Tarea,res.Error)
		}else{
			fmt.Printf("La %s se %s con un tiempo de %v\n",res.Tarea,res.Resultado,res.Tiempo)
		}
	}
}
