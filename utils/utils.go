package utils

import "fmt"

func ContarNumerosPares() int {
	contador := 0
	fmt.Println("\n=== Números del 1 al 10 ===")

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