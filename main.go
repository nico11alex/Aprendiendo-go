package main

import (
	"fmt"
)

type error interface {
	Error() string
}

type ErrorNumeros struct {
	Valor   int
	Mensaje string
}

func (eN *ErrorNumeros) Error() string {
	switch {
	case eN.Valor == 0:
		return fmt.Sprintf("Lo siento pero el segundo valor no puede ser %d", eN.Valor)
	default:
		return fmt.Sprintln("Error")
	}
}

func division(a int, b int) (resultado int, err error) {
	if b == 0 {
		errores := ErrorNumeros{
			Valor: b,
		}
		errores.Mensaje = errores.Error()
		return 0, fmt.Errorf("Error: %s", errores.Mensaje)
	} else {
		resultado := a + b
		return resultado, nil
	}
}

func main() {
	var a int = 1
	var b int = 0

	suma, err := division(a, b)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Se cierra el programa")
		return
	}
	fmt.Println("Todo salio bien")
	fmt.Println(suma)

}
