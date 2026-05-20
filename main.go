package main

import (
	"fmt"
)

type ErrorUserValidation struct {
	Campo   string
	Mensaje string
}

type user struct {
	generalInformation generalInformation
	securityData       securityData
	actionsUser        actionsUser
}

type generalInformation struct {
	Name  string
	Age   int
	Email string
}

type securityData struct {
	password string
}

type actionsUser struct {
	active bool
}

func (e ErrorUserValidation) Error() string {
	return fmt.Sprintf("error en %s\nproblema: %s", e.Campo, e.Mensaje)
}

func (u *user) Active() {
	u.actionsUser.active = true
}

func validateData(name string, age int, email string, password string) error {
	if err := validateUserName(name); err != nil {
		return err
	}
	if err := validateAge(age); err != nil {
		return err
	}
	if err := validatePassword(password); err != nil {
		return err
	}
	if err := validateEmail(email); err != nil {
		return  err
	}

	return nil
}

func createUser(name string, age int, email string, password string) (user, error) {
	if err:=validateData(name,age,email,password);err != nil{
		return user{},err
	}

	user := user{
		generalInformation: generalInformation{
			Name:  name,
			Age:   age,
			Email: email,
		},
		securityData: securityData{
			password: password,
		},
	}
	user.Active()
	return user, nil
}

func validateUserName(userName string) error {
	if len(userName) < 3 {
		return ErrorUserValidation{
			Campo:   "nombre de usuario",
			Mensaje: "el nombre es muy corto",
		}
	}
	return nil
}

func validateAge(edad int) error {
	if edad < 18 {
		return ErrorUserValidation{
			Campo:   "edad",
			Mensaje: "eres menor de edad y no podemos validar tu usuario",
		}
	}
	return nil
}

func validatePassword(password string) error {
	if len(password) < 8 {
		return ErrorUserValidation{
			Campo:   "contraseña",
			Mensaje: "la contraseña es muy corta",
		}
	}
	return nil
}

func validateAdmin(userName string, password string) bool {
	return userName == "admin" && password == "admin123"
}

func validateEmail(email string) error {
	for _, char := range email {
		if char == '@' {
			return nil
		}
	}
	return ErrorUserValidation{
		Campo:   "email",
		Mensaje: "no tiene formato correcto",
	}
}

func main() {
	resultado, err := createUser("admin", 19, "admin1222@","xskaxksa")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Usuario creado correctamente")
	fmt.Printf("Nombre: %s\nEdad: %d\nEmail: %s",resultado.generalInformation.Name,resultado.generalInformation.Age,resultado.generalInformation.Email)
	
}
