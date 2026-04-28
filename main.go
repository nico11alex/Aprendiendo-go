package main

import "fmt"

type Vehiculo struct {
	Marca   string
	Modelo string
	Año int
}


func main() {
	vehiculo := Vehiculo{
		Marca: "Ford",
		Modelo: "Mustang gtr",
		Año: 2026,
	}

	vehiculo2 := Vehiculo{} 

	vehiculo2.Marca = "Mazda"
	vehiculo2.Modelo= "CX30"
	vehiculo2.Año = 2023

	fmt.Println(vehiculo.Marca,vehiculo.Año)
	fmt.Println(vehiculo2.Modelo)


}