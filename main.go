package main

import (
	"errors"
	"fmt"
)

type ErrorRegistro struct {
	Codigo  int
	Campo   string
	Mensaje string
}

func (e ErrorRegistro) Error() string {
	return fmt.Sprintf("Error en el campo %s: %s", e.Campo, e.Mensaje)
}

func registrarUsuario(nombre string, edad int, email string, password string) (string, error) {
	var errs []error
	var err error
	if err = validarNombre(nombre); err != nil {
		errs = append(errs, err)
	}
	if err = validarEdad(edad); err != nil {
		errs = append(errs, err)
	}
	if err = validarEmail(email); err != nil {
		errs = append(errs, err)
	}

	if err = validarPassword(password); err != nil {
		errs = append(errs, err)
	}

	if len(errs) > 0 {
		return "", errors.Join(errs...)
	}

	return "Todo correcto", nil

}

func validarNombre(nombre string) error {
	if len(nombre) < 3 {
		return ErrorRegistro{
			Codigo:  1,
			Campo:   "nombre",
			Mensaje: "El nombre es muy corto",
		}
	}
	return nil
}

func validarEdad(edad int) error {
	if edad < 18 || edad > 100 {
		return ErrorRegistro{
			Codigo:  2,
			Campo:   "edad",
			Mensaje: "No estas en el rango de edad",
		}
	}
	return nil
}

func validarEmail(email string) error {
	for _, char := range email {
		if char == '@' {
			return nil
		}
	}
	return ErrorRegistro{
		Codigo:  3,
		Campo:   "email",
		Mensaje: "no tiene un formato correcto",
	}
}

func validarPassword(password string) error {
	var errs []error
	var err error
	var letraMayuscula bool
	var numero bool
	if len(password) > 0 {
		err = ErrorRegistro{
			Codigo:  4,
			Campo:   "contraseña",
			Mensaje: "muy corta",
		}
		errs = append(errs, err)
	}
	for _, char := range password {
		if char >= 'A' && char <= 'Z' {
			letraMayuscula = true
		}
		if char >= '0' && char <= '9' {
			numero = true
		}
	}
	if !letraMayuscula {
		err = ErrorRegistro{
			Codigo:  5,
			Campo:   "contraseña",
			Mensaje: "Le falta una letra mayuscula",
		}
		errs = append(errs, err)
	}
	if !numero {
		err = ErrorRegistro{
			Codigo:  6,
			Campo:   "contraseña",
			Mensaje: "Le falta un número",
		}
		errs = append(errs, err)
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}
	return nil
}

func main() {
	pruebas := []struct {
		nombre   string
		edad     int
		email    string
		password string
	}{
		{"Nicolas", 25, "nicolas@gmail.com", "12234234"}, // éxito
		{"Ni", 20, "nicolasgmail.com", "Nicolas12"},      // nombre corto + sin @
		{"Carlos", 15, "carlos@ok.com", "sadsadsajdk23"}, // edad inválida
		{"Ana", 5, "ana", "as"},                          // múltiples errores
	}

	for _, p := range pruebas {
		fmt.Printf("Registrando: %s (%d) - %s, contraseña: %s\n", p.nombre, p.edad, p.email, p.password)

		_, err := registrarUsuario(p.nombre, p.edad, p.email, p.password)

		if err != nil {
			fmt.Printf("❌ Errores: %v\n\n", err)
		} else {
			fmt.Println("✅ Registro exitoso")
		}
	}
}
