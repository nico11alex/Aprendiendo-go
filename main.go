package main

import (
	"errors"
	"fmt"
)

type Comprable interface {
	Precio() int
}

type Mazda struct {
	Modelo int
	Valor  int
}

type ErrorPresupuesto struct{
	Diferencia int
}

func (m Mazda) Precio() int {
	return m.Valor
}

func RealizarCompra(c Comprable, presupuesto int,e *ErrorPresupuesto) (string, error) {
	if c.Precio() > presupuesto {
		Diferencia(c,presupuesto,e) 
		MensajeError := fmt.Sprintf("Te faltan solo %d meses de ahorro si guardas 2 millones al mes.", e.Diferencia)
		err := errors.New(MensajeError)
		return "", err
	} else {
		return "Felicidades puedes comprar el auto", nil
	}
}

func Diferencia(c Comprable, presupuesto int, e *ErrorPresupuesto){
	for c.Precio() > presupuesto {
		presupuesto += 2000000
		e.Diferencia ++
	}
}

func main() {
	mazda := Mazda{Modelo: 2026, Valor: 140000000}
	errores := ErrorPresupuesto{Diferencia: 0}
	mensaje, errorCompra := RealizarCompra(mazda, 13000000,&errores)
	if errorCompra != nil{
		fmt.Println(errorCompra)
	}else{
		fmt.Println(mensaje)
	}
}
