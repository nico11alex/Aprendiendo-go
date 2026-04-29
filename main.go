package main

import "fmt"

type Figura interface{
	Area() int
}

type Rectangulo struct{
	Ancho int
	Alto int
}

type Cuadrado struct{
	Lado int
}

func ImprimirInformacion(f Figura) {
    fmt.Printf("El área de esta figura es: %d\n", f.Area())
}

func(r Rectangulo) Area()int{
	return r.Alto*r.Ancho
}

func(c Cuadrado) Area()int{
	return c.Lado*c.Lado
}

func(r *Rectangulo) Escalar(factor int){
	if(factor <=0){
		return
	}
	r.Alto *= factor
	r.Ancho *= factor

}

func(c *Cuadrado) Escalar(factor int){
	if(factor <=0){
		return
	}
	c.Lado *= factor

}

func main() {
	rectangulo := Rectangulo{
		Ancho: 4,
		Alto: 3,
	}
	cuadrado := Cuadrado{
		Lado: 5,
	}

	// rectangulo

	fmt.Printf("El area del rectangulo es de %d\n",rectangulo.Area())
	rectangulo.Escalar(2)
	fmt.Printf("El area del rectangulo es de %d\n",rectangulo.Area())

	// cuadrado

	fmt.Printf("El area del cuadrado es de %d\n",cuadrado.Area())
	cuadrado.Escalar(2)
	fmt.Printf("El area del cuadrado es de %d\n",cuadrado.Area())

	// Figura
	ImprimirInformacion(rectangulo)
	ImprimirInformacion(cuadrado)
}