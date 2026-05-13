package main

import (
	"fmt"
)

type ErrorValidacion struct {
	Campo   string
	Mensaje string
}

func (e ErrorValidacion) Error() string {
	return fmt.Sprintf("%s: %s", e.Campo, e.Mensaje)
}

func crearUsuario(username, email string)(string,error){
	if len(username) < 4 {
		return "",ErrorValidacion{
			Campo: "username",
			Mensaje: fmt.Sprintf("el nombre %s es muy corto para ser creado",username),
		}
	}
	formatoEmailValido := false
	for _,char := range email{
		if char == '@'{
			formatoEmailValido = true
		}
	}
	
	if !formatoEmailValido{
		return "",ErrorValidacion{
			Campo: "email",
			Mensaje: fmt.Sprintf("el email %s no es valido por que no contiene @.",email),
		}
	}
	return "Todo correcto",nil
}

func main() {
	username := "Nic"
	email := "asdwsrwfw"

	resultado, err := crearUsuario(username,email)
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(resultado)

}
