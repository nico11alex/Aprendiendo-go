package main

import "fmt"

type Imprimible interface {
	Info() string
}

type Libro struct {
	Titulo string
	Autor  string
}

type Pelicula struct {
	Titulo   string
	Director string
}

func (l Libro) Info() string {
	return fmt.Sprintf("Libro: %s de %s",l.Titulo,l.Autor)
}

func (p Pelicula) Info() string {
	return fmt.Sprintf("Pelicula: %s dirigida por %s",p.Titulo,p.Director)
}

func main() {
	libreria := []Imprimible{
		Libro{Titulo: "Habitos Atomicos", Autor: "James Clear"},
		Pelicula{Titulo: "Juego de honor", Director: "Tomas Carter"},
	}
	for _, contenido := range libreria {
		fmt.Printf("%s\n",contenido.Info())
	}
}
