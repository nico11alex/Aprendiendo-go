package main

import (
	"fmt"
	"strconv"
)

func ejercicio(mensualidad int) string {
	precioMazda := 126000000
	var ahorro uint32
	for {
		precioFaltante := strconv.Itoa(precioMazda - int(ahorro))
		ahorro += (uint32(mensualidad) * 30) / 100
		if ahorro > uint32(precioMazda) {
			return "¡Felicidades, le compraste el carro a tu papá!"
		} else {
			fmt.Println("Te faltan " + precioFaltante + " pesos para comprarle el mazda a tu papá")
		}

	}
}
