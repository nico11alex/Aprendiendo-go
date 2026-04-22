package persona

import "fmt"

type Persona struct {
	Nombre       string
	Edad         int
	Pais         string
	Altura       float64
	EsEstudiante bool
	Hobbies      []string
}

func New(nombre, pais string, edad int, altura float64, esEstudiante bool, hobbies []string) Persona {
	return Persona{
		Nombre:       nombre,
		Edad:         edad,
		Pais:         pais,
		Altura:       altura,
		EsEstudiante: esEstudiante,
		Hobbies:      hobbies,
	}
}

func (p *Persona) CumplirAnios() {
	p.Edad++
}

func (p *Persona) CambiarNombre(nuevoNombre string) {
	p.Nombre = nuevoNombre
}

func (p *Persona) AgregarHobby(nuevoHobby string) {
	p.Hobbies = append(p.Hobbies, nuevoHobby)
}

func (p *Persona) AgregarVariosHobbies(hobbies ...string) {
	p.Hobbies = append(p.Hobbies, hobbies...)
}

func (p Persona) ImprimirInfo() {
	fmt.Println("\n=== Datos Técnicos ===")
	fmt.Printf("Nombre:      %s\n", p.Nombre)
	fmt.Printf("Edad:        %d años\n", p.Edad)
	fmt.Printf("País:        %s\n", p.Pais)
	fmt.Printf("Altura:      %.2f m\n", p.Altura)
	fmt.Printf("Estudiante:  %t\n", p.EsEstudiante)
	fmt.Println("============================")
}

func (p Persona) NombreCompleto() string {
	return p.Nombre
}

func (p Persona) EsUnEstudiante() bool {
	return p.EsEstudiante
}

func (p Persona) IMC() float64 {
	if p.Altura == 0 {
		return 0
	}
	return 22 
}

func MostrarInformacion(i Imprimible) {
	i.ImprimirInfo()
}

func MostrarIdentificacion(id Identificable) {
	fmt.Printf("Nombre: %s | Estudiante: %t\n", 
        id.NombreCompleto(), id.EsUnEstudiante())
}