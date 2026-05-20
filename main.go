package main

import (
	"fmt"
)

type ErrorUserValidation struct {
	Campo  string
	Mensaje string
}

func (e ErrorUserValidation) Error() string {
	return fmt.Sprintf("error en %s\nproblema: %s", e.Campo, e.Mensaje)
}

func userValidation(userName string, edad int, password string) (string, error) {
	if err := validateUserName(userName); err != nil {
		return "", err
	}
	if err := validateEdad(edad); err != nil {
		return "", err
	}
	if err := validatePassword(password); err != nil {
		return "", err
	}
	if validateAdmin(userName, password) {
		return "Bienvenido admin", nil
	}
	return "usuario validado de manera correcto", nil
}

func validateUserName(userName string) error {
	if len(userName) < 3 {
		return ErrorUserValidation{
			Campo:  "nombre de usuario",
			Mensaje: "el nombre es muy corto",
		}
	}
	return nil
}

func validateEdad(edad int) error {
	if edad < 18 {
		return ErrorUserValidation{
			Campo:  "edad",
			Mensaje: "eres menor de edad y no podemos validar tu usuario",
		}
	}
	return nil
}

func validatePassword(password string) error {
	if len(password) < 8 {
		return ErrorUserValidation{
			Campo:  "contraseña",
			Mensaje: "la contraseña es muy corta",
		}
	}
	return nil
}

func validateAdmin(userName string, password string) bool {
	return userName == "admin" && password == "admin123"
}

func main() {
	resultado,err := userValidation("admin",19,"admin1222")
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(resultado)
}
