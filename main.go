package main

import "fmt"

func main() {
	nombre := "Nicolas"
	edad := 18
	pais := "Colombia"

	altura := 1.80
	esEstudiante := true

	fmt.Printf("Nombre:%s \nEdad:%d años \nPais:%s \nAltura:%.2f \nEstudiante:%t\n", nombre, edad, pais, altura, esEstudiante)
	fmt.Println("Estoy aprendiendo Go en Windows")
}
