package main

import "fmt"

type Persona struct{
	Nombre string
	Edad int
	Pais string
	Altura float64
	EsEstudiante bool
	Hobbies []string
}

const ( 
		mayorDeEdad           = 18
		edadMinimaAdolescente = 13 
	)

func imprimirInfoPersonal(nombre string ,edad int,pais string,altura float64,esEstudiante bool){
	fmt.Println("\n=== Datos Técnicos ===")
	fmt.Printf("Nombre:      %s\n", nombre)
	fmt.Printf("Edad:        %d años\n", edad)
	fmt.Printf("País:        %s\n", pais)
	fmt.Printf("Altura:      %.2f m\n", altura)
	fmt.Printf("Estudiante:  %t\n", esEstudiante)
	fmt.Println("============================")
}

func determinarCategoriaEdad(edad,mayorDeEdad,edadMinimaAdolescente int)string{
	switch{
	case edad < 0 || edad > 130:
		return "Tu edad no puede ser cierta"
	case edad >= mayorDeEdad:
		return  "Eres mayor de edad"
	case edad >= edadMinimaAdolescente:
		return  "Eres adolescente"
	default:
		return "Eres un niño"
	}

}

func contarNumerosPares()int{
	contador :=0
	for i := 1; i <= 10; i++ {
		if i%2 == 0{
			fmt.Printf("El número %d es par\n",i)
			contador ++
		}else{
			fmt.Printf("El número %d es impar\n",i)
		}
	}

	return contador
}

func imprimirHobbies(hobbies []string){
	fmt.Println("Tus hobbies son:")
	for indice,hobby := range hobbies{
		fmt.Printf("Tu hobby número %d es %s\n",indice+1, hobby)
	}
}

func agregarHobby(hobbies []string, nuevoHobby string) []string{
	return append(hobbies, nuevoHobby)
}

func nuevaPersona(nombre, pais string, edad int, altura float64, esEstudiante bool, hobbies []string) Persona{
	persona := Persona{
		Nombre: nombre,
		Edad: edad,
		Pais: pais,
		Altura: altura,
		EsEstudiante: esEstudiante,
		Hobbies: hobbies,
	}

	return persona
}

func imprimirPersona(persona Persona){
	fmt.Println("\n=== Datos Técnicos ===")
	fmt.Printf("Nombre:      %s\n", persona.Nombre)
	fmt.Printf("Edad:        %d años\n", persona.Edad)
	fmt.Printf("País:        %s\n", persona.Pais)
	fmt.Printf("Altura:      %.2f m\n", persona.Altura)
	fmt.Printf("Estudiante:  %t\n", persona.EsEstudiante)
	imprimirHobbies(persona.Hobbies)
	fmt.Println("============================")
}

func main() {

	hobbies:=[]string {"Programar","Leer","Baloncesto","Video juegos","Escribir"}

	nombre := "Nicolas"
	edad := 18
	pais := "Colombia"
	altura := 1.80
	esEstudiante := true
	
	hobbies = agregarHobby(hobbies,"Comer")

	yo := nuevaPersona(nombre,pais,edad,altura,esEstudiante,hobbies)
	imprimirPersona(yo)

	fmt.Println("")
	fmt.Println(determinarCategoriaEdad(edad,mayorDeEdad,edadMinimaAdolescente))
	fmt.Println("")

	fmt.Println("Números del 1 al 10")
	fmt.Printf("Hemos encontrado %d números pares\n",contarNumerosPares())
	fmt.Println("")
	fmt.Println("Estoy aprendiendo Go en Windows")
}
