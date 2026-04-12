package main

import "fmt"

type Carro struct {
	Nombre string
	Modelo int
	Precio int
}

type Persona struct {
	Nombre string
	Edad   int
	Sueldo int
}

func listaDeCarros(carros []Carro) {
	fmt.Println("Nuestros carros")
	for _, carro := range carros {
		fmt.Printf("%s ", carro.Nombre)
	}
	fmt.Println("")
	fmt.Printf("Escoja un carro:")
}

func validacionEdad(persona Persona)bool{
	if persona.Edad < 18{
		return false
	}
	return true
}

func calcularCuantoTiempoFaltaEdad(persona Persona)int{
	tiempoFaltante := 18 - persona.Edad
	return tiempoFaltante
}

func confirmarCantidadDinero(nombreCarro,nombrePersona string, modelo, precio,sueldo int)bool{
	if(sueldo>=precio){
		fmt.Printf("Felicidades %s puedes comprar un %s modelo %d con tu sueldo",nombrePersona,nombreCarro,modelo)
		return true
	}
	return false
}

func calcularCuantoTiempoFaltaDinero(precio,sueldo int)int{
	meses := 0
	ahorro:=0
	for ahorro <= precio {
		ahorro += sueldo
		meses ++
	}
	return meses
}

func main() {

	var carroElejido string
	var nombre string
	var edad int
	var sueldo int

	fmt.Printf("Cual es tu nombre: ")
	_, err := fmt.Scanln(&nombre)
	if err != nil {
		fmt.Println("Tu nombre tiene un formato incorrecto")
		return
	}

	fmt.Printf("Cual es tu edad: ")
	_, err = fmt.Scanln(&edad)
	if err != nil {
		fmt.Println("Tu edad no puede ser guardada revisala")
		return
	}

	fmt.Printf("Cual es tu sueldo: ")
	_, err = fmt.Scanln(&sueldo)
	if err != nil {
		fmt.Println("Tu sueldo no puede ser guardada revisala")
		return
	}

	persona := Persona{
		nombre,
		edad,
		sueldo,
	}

	carros := []Carro{
		{"Mustang", 2020, 200000000},
		{"Mazda", 2021, 15000000},
	}

	if !validacionEdad(persona){
		fmt.Printf("Lo siento %s pero aun te faltan %d años para tratar de comprar un carro",persona.Nombre,calcularCuantoTiempoFaltaEdad(persona))
		return
	}

	listaDeCarros(carros)

	_, err = fmt.Scanln(&carroElejido)
	if err != nil {
		fmt.Println("Eso no es un carro")
		return
	}

	for _, carro := range carros {
		if carro.Nombre == carroElejido {
			if confirmarCantidadDinero(carro.Nombre,persona.Nombre,carro.Modelo,carro.Precio,persona.Sueldo){
				return
			}

			fmt.Printf("Perdon %s aun te falta %d meses para comprar el %s que quieres",persona.Nombre,calcularCuantoTiempoFaltaDinero(carro.Precio,sueldo),carro.Nombre)
			return
		}
	}

	fmt.Println("Carro no encontrado")
}
