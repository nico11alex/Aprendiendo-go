package main

import (
	"fmt"
)

func validarPassword(contraseña string) (bool, error) {
	if len(contraseña) < 8 {
		return false, fmt.Errorf("perdon pero la contraseña es muy corta")
	}

	tieneNumero := false
	tieneLetraMayuscula := false

	for _, char := range contraseña {
		if char >= '0' && char <= '9' {
			tieneNumero = true
		}
		if char >= 'A' && char <= 'Z' {
			tieneLetraMayuscula = true
		}
	}

	if !tieneNumero {
		return false, fmt.Errorf("perdon pero la contraseña debe tener un numero")
	}
	if !tieneLetraMayuscula {
		return false, fmt.Errorf("perdon pero la contraseña debe tener una letra en mayuscula")
	}
	return true, nil
}

func main() {
	contraseña := "Nicolas1"
	resultado, err := validarPassword(contraseña)
	if err != nil {
		fmt.Println(err)
		return
	}
	if resultado {
		fmt.Println("Contraseña valida")
	}

}
