package main

import (
	"encoding/json"
	"net/http"
)

type Persona struct {
	Nombre string `json:"nombre"`
	Edad   int    `json:"edad"`
}

type Carro struct {
	Nombre string `json:"nombre"`
	Precio int `json:"precio"`
}

func obtenerPersona(w http.ResponseWriter, r *http.Request) {

	personas := []Persona{
		{"Nicolas", 20},
		{"Sara", 22},
		{"Jesus", 15},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(personas)
}

func obtenerCarros(w http.ResponseWriter, r *http.Request){

	carros :=[]Carro{
		{"Mazda", 114350000},
		{"Mustang gt", 279990000},
	}

	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(carros)
}

func main() {
	http.HandleFunc("/persona", obtenerPersona)
	http.HandleFunc("/carros",obtenerCarros)
	http.ListenAndServe(":8080", nil)
}