package main

import "fmt"

func main() {
	const mayorDeEdad = 18
	const edadMinimaAdolescente = 13

	nombre := "Nicolas"
	edad := 18
	pais := "Colombia"
	altura := 1.80
	esEstudiante := true

	if edad < 0 || edad > 130{
		fmt.Printf("%s tu edad no puede ser cierta\n",nombre)
	}else if edad >= mayorDeEdad{
		fmt.Printf("%s eres mayor de edad\n",nombre)
	}else if edad >= edadMinimaAdolescente{
		fmt.Printf("%s eres adolecente\n",nombre)
	}else{
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

	fmt.Printf("Nombre:%s \nEdad:%d años \nPais:%s \nAltura:%.2f \nEstudiante:%t\n", nombre, edad, pais, altura, esEstudiante)
	fmt.Println("Estoy aprendiendo Go en Windows")
}
