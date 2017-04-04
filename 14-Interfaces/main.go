package main

import "fmt"

/*
Intefaces -> ayudan a estandarizar alguna estructa, obligando a implentar funciones establecidas en ellas

SINTAXIS
	type Nombre interface{
		funcion1() tipoRetorno
		funcion2()
	}

*/

//Persona estrucutra para Persona
type Persona struct {
	Nombre, Pais string
}

func (p Persona) getNombre() string {
	return p.Nombre
}

func (p Persona) getPais() string {
	return p.Pais
}

//Presentarse func
func Presentarse(u Usuario) {
	fmt.Println("Hola soy:", u.getNombre(), "y mi pais es:", u.getPais())
}

//Usuario Interfaz para la cual se debe implementar 2 funciones getNombre y getPais
type Usuario interface {
	getNombre() string
	getPais() string
}

func main() {
	andres := Persona{"ANDRES LATORRE", "COLOMBIA"}
	Presentarse(andres)
}
