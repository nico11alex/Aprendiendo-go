package main

import "fmt"

func calcularPromedio(numeros []float64) (float64, error) {
	if len(numeros) == 0 {
		return 0.0, fmt.Errorf("Error no se puede calcular el promedio de una lista vacía")
	}

	var suma float64
	for _, numero := range numeros {
		suma += numero
	}
	
	promedio := suma / float64(len(numeros))
	return promedio, nil
}

func main() {
	numeros := []float64{}

	resultado, err := calcularPromedio(numeros)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resultado)
}
