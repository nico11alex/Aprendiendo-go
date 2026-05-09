package main

import "fmt"

func registrarUsuario(nombre string, edad int) (string, error) {
	err1 := validarEdad(edad)
	err2 := validarNombre(nombre)
	switch {
	case err1 != nil && err2 != nil:
		return "", fmt.Errorf("Errores: 1.%w,\n2.%w", err1, err2)
	case err1 != nil:
		return "", fmt.Errorf("Error: %w", err1)
	case err2 != nil:
		return "", fmt.Errorf("Error: %w", err2)
	default:
		return "Registro exitoso", nil
	}
}

func validarEdad(edad int) error {
	if edad < 18 || edad > 100 {
		return fmt.Errorf("Estas fuera de el rango de edad")
	}
	return nil
}

func validarNombre(nombre string) error {
	if len(nombre) < 3 {
		return fmt.Errorf("Perdon pero el nombre esta muy corto")
	}
	return nil
}

func main() {
	resultado, err := registrarUsuario("Nicolas", 18)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resultado)
}
