package main

import "fmt"

type Rectangulo struct{
	Ancho int
	Alto int
}

func(r Rectangulo) Area()int{
	return r.Alto*r.Ancho
}

func(r *Rectangulo) Escalar(factor int){
	if(factor <=0){
		return
	}
	r.Alto *= factor
	r.Ancho *= factor

}

func main() {
	rectangulo := Rectangulo{
		Ancho: 4,
		Alto: 3,
	}

	fmt.Printf("El area del rectangulo es de %d\n",rectangulo.Area())
	rectangulo.Escalar(2)
	fmt.Printf("El area del rectangulo es de %d\n",rectangulo.Area())
}