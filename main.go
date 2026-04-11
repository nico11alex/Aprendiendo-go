package main

import (
	"fmt"
	"net/http"
)

func saludo(w http.ResponseWriter, r *http.Request) {

	nombre := r.URL.Query().Get("nombre")

	if nombre == "" {
		nombre = "Invitado"
	}

	fmt.Fprintf(w, "Hola %s bienvenido a mi primera API en go\n", nombre)
	fmt.Fprintf(w, "El metodo request utilizado es %s\n", r.Method)
	fmt.Fprintf(w, "La ruta utilizada es %s\n", r.URL)
}

func main() {

	http.HandleFunc("/saludo", saludo)
	fmt.Println("Servidor en http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}
