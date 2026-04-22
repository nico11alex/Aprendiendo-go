package main

import (
	"errors"
	"fmt"
)

type Persona struct {
	Nombre       string
	Edad         int
	Pais         string
	Altura       float64
	EsEstudiante bool
	Hobbies      []string
}

type Imprimible interface {
	ImprimirInfo()
}

type Identificable interface {
	NombreCompleto() string
	EsUnEstudiante() bool
}

func determinarCategoriaEdad(edad int) (string, error) {
	switch {
	case edad < 0 || edad > 130:
		return "", errors.New("edad fuera de rango válido")
	case edad >= 18:
		return "Eres mayor de edad", nil
	case edad >= 13:
		return "Eres adolescente", nil
	default:
		return "Eres un niño", nil
	}

}

func MostrarInformacion(i Imprimible) {
	i.ImprimirInfo()
}

func MostrarIdentificacion(id Identificable) {
	fmt.Printf("Nombre: %s | Estudiante: %t\n", 
        id.NombreCompleto(), id.EsUnEstudiante())
}

func contarNumerosPares() int {
	contador := 0
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Printf("El número %d es par\n", i)
			contador++
		} else {
			fmt.Printf("El número %d es impar\n", i)
		}
	}

	return contador
}

func imprimirHobbies(hobbies []string) {
	fmt.Println("Tus hobbies son:")
	for indice, hobby := range hobbies {
		fmt.Printf("Tu hobby número %d es %s\n", indice+1, hobby)
	}
}

func nuevaPersona(nombre, pais string, edad int, altura float64, esEstudiante bool, hobbies []string) Persona {
	persona := Persona{
		Nombre:       nombre,
		Edad:         edad,
		Pais:         pais,
		Altura:       altura,
		EsEstudiante: esEstudiante,
		Hobbies:      hobbies,
	}

	return persona
}

func (persona Persona) ImprimirInfo() {
	fmt.Println("\n=== Datos Técnicos ===")
	fmt.Printf("Nombre:      %s\n", persona.Nombre)
	fmt.Printf("Edad:        %d años\n", persona.Edad)
	resultado, err := determinarCategoriaEdad(persona.Edad)
	if err != nil {
		fmt.Println("Error :", err)
	}
	fmt.Printf("Categoría:   %s\n", resultado)
	fmt.Printf("País:        %s\n", persona.Pais)
	fmt.Printf("Altura:      %.2f m\n", persona.Altura)
	fmt.Printf("Estudiante:  %t\n", persona.EsEstudiante)
	imprimirHobbies(persona.Hobbies)
	fmt.Println("============================")
}

func (p *Persona) CumplirAnios() {
	p.Edad++
}

func (p *Persona) CumplirAniosYFelicitar() {
	p.Edad++
	fmt.Printf("¡Feliz cumpleaños, %s! Ahora tienes %d años\n", p.Nombre, p.Edad)
}

func (p *Persona) CambiarNombre(nuevoNombre string) {
	p.Nombre = nuevoNombre
}

func (p *Persona) AgregarHobby(nuevoHobby string) {
	p.Hobbies = append(p.Hobbies, nuevoHobby)
}

func (p Persona) NombreCompleto() string {
	return p.Nombre
}

func (p Persona) EsUnEstudiante() bool {
	return p.EsEstudiante
}

func (p Persona) IMC() float64 {
	peso := (p.Altura * p.Altura) * 22
	return peso / (p.Altura * p.Altura)
}

func (p *Persona) AgregarVariosHobbies(hobbies ...string) {
	p.Hobbies = append(p.Hobbies, hobbies...)
}

func (p *Persona) InfoResumida() (string, int, []string) {
	return p.Nombre, p.Edad, p.Hobbies
}

func (p Persona) EsMayorDeEdadEnPais(pais string) bool {
	if p.Edad >= 18 && pais == p.Pais {
		return true
	}
	return false
}

func main() {

	hobbies := []string{"Programar", "Leer", "Baloncesto", "Video juegos", "Escribir"}

	nombre := "Nicolas"
	edad := 18
	pais := "Colombia"
	altura := 1.80
	esEstudiante := true

	yo := nuevaPersona(nombre, pais, edad, altura, esEstudiante, hobbies)

	yo.CumplirAnios()
	yo.CambiarNombre("Jesus")
	yo.AgregarHobby("Correr")
	yo.AgregarVariosHobbies("Viajar", "Gimnasio", "Musica")
	yo.ImprimirInfo()

	fmt.Println("Nombre completo:", yo.NombreCompleto())
	fmt.Println("¿Es estudiante?", yo.EsUnEstudiante())
	fmt.Printf("IMC aproximado: %.2f\n", yo.IMC())

	yo.CumplirAniosYFelicitar()

	nombre, edad, hobbies = yo.InfoResumida()
	fmt.Printf("Resumen: %s, %d años, %d hobbies\n", nombre, edad, len(hobbies))

	fmt.Println("¿Es mayor de edad en Colombia?", yo.EsMayorDeEdadEnPais("Colombia"))
	fmt.Println("¿Es mayor de edad en México?", yo.EsMayorDeEdadEnPais("México"))

	MostrarInformacion(yo)
	MostrarIdentificacion(yo)

	fmt.Println("")
	fmt.Println("Números del 1 al 10")
	fmt.Printf("Hemos encontrado %d números pares\n", contarNumerosPares())
	fmt.Println("")
	fmt.Println("Estoy aprendiendo Go en Windows")
}
