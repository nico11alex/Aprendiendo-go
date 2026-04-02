package main

import "fmt"

type Persona struct {
	Nombre string
	Edad   int
	Dinero int
}

func esMayorDeEdad(edad int)bool{
	return edad >=18
}

func puedeComprar(dinero,precio int)bool{
	return dinero >= precio
}


func main() {

	var nombre string
	var edad int
	var dinero int

	fmt.Println("Escriba su nombre: ")
	fmt.Scan(&nombre)

	fmt.Println("Escriba su edad: ")
	_, err := fmt.Scanln(&edad)

	if err != nil {
		fmt.Println("Error: ingresa un número válido")
		return
	}

	fmt.Println("Escriba su dinero ahorrado: ")
	_,err = fmt.Scanln(&dinero)

	if err != nil {
		fmt.Println("Error: ingresa un número válido")
		return
	}

	precio:=2000000

	personas := Persona{
		Nombre: nombre,
		Edad: edad,
		Dinero: dinero,
	}

	if !esMayorDeEdad(personas.Edad){
		fmt.Printf("%s perdon pero no puedes comprar un carro porque tienes %d años y eres menor de edad\n",personas.Nombre,personas.Edad)
		return
	}

	if !puedeComprar(personas.Dinero,precio){
		fmt.Printf("%s perdon pero con tu dinero no te alcanza el carro\n",personas.Nombre)
		return
	}
	
	fmt.Printf("%s felicidades puedes comprar tu Mustang de 2000000",personas.Nombre)


}
