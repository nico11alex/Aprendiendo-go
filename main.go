package main

import "fmt"
func puedeComprar(dineroAhorrado, precio int)bool{
	return dineroAhorrado >= precio
}

func main() {
	nombre := "Nicolas"
	dineroAhorrado := 3000000
	carros := []string{"Mustang", "Camaro", "Raptor"}
	precios := []int{2000000, 100000, 2500000}

	for i,carro := range carros{
		precio := precios[i]
		if puedeComprar(dineroAhorrado,precio){
			fmt.Printf("Felicidades %s tu tienes $%d y puedes comprar un %s por que cuesta $%d\n", nombre, dineroAhorrado, carro, precio)
		}else {
			fmt.Printf("%s NO puedes comprar el %s ($%d)\n", nombre, carro, precio)
		}
	}

}
