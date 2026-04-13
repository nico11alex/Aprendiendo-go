package main

import "fmt"

func main() {
	const(mayorDeEdad = 18
	edadMinimaAdolescente = 13)

	nombre := "Nicolas"
	edad := 18
	pais := "Colombia"
	altura := 1.80
	esEstudiante := true
	var contador int

	switch{
	case edad < 0 || edad > 130:
		fmt.Printf("%s tu edad no puede ser cierta\n",nombre)
	case edad >= mayorDeEdad:
		fmt.Printf("%s eres mayor de edad\n",nombre)
	case edad >= edadMinimaAdolescente:
		fmt.Printf("%s eres adolecente\n",nombre)
	default:
		fmt.Printf("%s eres un niño\n",nombre)
	}

	if altura <= 0 || altura > 3.5{
		fmt.Printf("%s tu altura parece no ser cierta\n",nombre)
	}else if altura > 1.90{
		fmt.Printf("%s eres muy alto\n",nombre)
	}else if altura < 1.60{
		fmt.Printf("%s eres de estatura baja\n",nombre)
	}else{
		fmt.Printf("%s tienes una estatura promedio\n",nombre)
	}

	for i := 1; i <= 10; i++ {
		if i%2 == 0{
			fmt.Printf("El número %d es par\n",i)
			contador ++
		}else{
			fmt.Printf("El número %d es impar\n",i)
		}
	}

	fmt.Printf("He contado hasta 10 y encontré %d números pares\n",contador)

	fmt.Println("\n=== Datos Técnicos ===")
	fmt.Printf("Nombre:      %s\n", nombre)
	fmt.Printf("Edad:        %d años\n", edad)
	fmt.Printf("País:        %s\n", pais)
	fmt.Printf("Altura:      %.2f m\n", altura)
	fmt.Printf("Estudiante:  %t\n", esEstudiante)
	fmt.Println("Estoy aprendiendo Go en Windows")
	fmt.Println("============================")
}
