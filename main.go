package main

import "fmt"

type Estudiante struct {
	Nombre   string
	Edad     int
	Promedio float64
	Activo   bool
}

func main() {
	yo := Estudiante{
		Nombre:   "Nicolas",
		Edad:     18,
		Promedio: 8.9,
		Activo:   true,
	}

	fmt.Printf("El estudiante %s tiene %d años y un promedio de %f\n",yo.Nombre,yo.Edad,yo.Promedio)
	yo.Activo = false
	fmt.Println(yo.Activo)

}