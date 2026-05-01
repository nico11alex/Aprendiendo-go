package main

import (
    "fmt"
    "sync"
    "time"
)

func trabajador(id int, wg *sync.WaitGroup, resultados chan<- string) {
    defer wg.Done()

    fmt.Printf("Trabajador %d empezó\n", id)
    time.Sleep(time.Duration(id) * 500 * time.Millisecond)

    mensaje := fmt.Sprintf("Trabajador %d terminó su trabajo", id)
    
    resultados <- mensaje   // ← Enviamos el resultado por el channel
}

func main() {
    var wg sync.WaitGroup
    resultados := make(chan string, 3)   // Channel con buffer de 3

    for i := 1; i <= 3; i++ {
        wg.Add(1)
        go trabajador(i, &wg, resultados)
    }

    // Esperamos que todas terminen
    wg.Wait()

    // Recibimos los resultados
    fmt.Println("\nResultados recibidos:")
    for i := 1; i <= 3; i++ {
        fmt.Println("→", <-resultados)
    }

    fmt.Println("\n¡Programa finalizado correctamente!")
}