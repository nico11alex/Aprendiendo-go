package main

import (
	"fmt"
	"aprendiendo-go/persona"
	"aprendiendo-go/utils"
)

func main() {
	hobbies := []string{"Programar", "Leer", "Baloncesto", "Videojuegos", "Escribir"}

	yo := persona.New("Jesús", "Colombia", 19, 1.80, true, hobbies)

	yo.CumplirAnios()
	yo.CambiarNombre("Jesús López")
	yo.AgregarVariosHobbies("Viajar", "Gimnasio", "Música")

	yo.ImprimirInfo()

	fmt.Println("Nombre completo:", yo.NombreCompleto())
	fmt.Println("¿Es estudiante?", yo.EsUnEstudiante())

	// Uso de interfaces desde el paquete persona
	// (tendrás que mover MostrarInformacion y MostrarIdentificacion al paquete persona)
	
	utils.ContarNumerosPares()
	persona.MostrarIdentificacion(yo)
	persona.MostrarInformacion(yo)

	fmt.Println("\nEstoy aprendiendo Go con paquetes!")
}