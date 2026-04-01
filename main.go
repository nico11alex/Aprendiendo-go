package main

import "fmt"

type Persona struct {
	nombre string
	edad   int
	dinero int
}

func main() {

	precio := 2000000

	personas := []Persona{
		{"Nicolas", 18, 2000000},
		{"Sara", 22, 3000000},
		{"Jesus", 15, 1500000},
	}

	for _, persona := range personas {
		fmt.Printf("%s cuentas con %d a tus %d\n",persona.nombre,persona.dinero,persona.edad)
		if persona.dinero >= precio{
			fmt.Printf("Puedes comprar un carro de 2000000\n")
		}else{
			fmt.Printf("No puedes comprar un carro de 2000000\n")
		}
	}

}
