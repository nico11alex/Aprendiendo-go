package main

import (
	"fmt"
	"sync"
)

func concesionario1(name string, wg *sync.WaitGroup, resultados chan<- string) {
    defer wg.Done()

    mensaje := fmt.Sprintf("%s tu mustang cuesta $279.990.000 COP en el consecionario Autoland Colombia",name)

    resultados <- mensaje
}

func concesionario2(name string, wg *sync.WaitGroup, resultados chan<- string) {
    defer wg.Done()

    mensaje := fmt.Sprintf("%s tu mustang cuesta $249.900.000 COP en el consecionario El Roble Motor (Pereira/Medellín/Cali):.",name)

    resultados <- mensaje
}

func concesionario3(name string, wg *sync.WaitGroup, resultados chan<- string) {
    defer wg.Done()

    mensaje := fmt.Sprintf("%s tu mustang cuesta $364.900.000 COP en el consecionario Los Coches",name)

    resultados <- mensaje
}


func main(){
    var wg sync.WaitGroup
    
    resultados := make(chan string, 3)

    
    wg.Add(1)
    go concesionario1("Nicolas",&wg,resultados)
    wg.Add(1)
    go concesionario2("Nicolas",&wg,resultados)
    wg.Add(1)
    go concesionario3("Nicolas",&wg,resultados)
    
    wg.Wait()
    close(resultados)

    fmt.Println("\nResultados recibidos:")
    for mesage := range resultados {
        fmt.Println("→", mesage)
    }
    
}