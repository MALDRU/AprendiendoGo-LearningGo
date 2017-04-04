package main

import "fmt"

/*
type -> se utiliza para crear tipos propios

Sintaxis
	type nombre base
	Ej:
	type segundos int

//para crear una base propia, revisar 12-Estructuras
*/

type segundos int
type minutos int

func main() {
	var seg segundos = 10
	var min minutos = 5

	fmt.Println("Segundos:", seg, "\nMinutos:", min)

	//sumar minutos y segundos
	fmt.Println("Total en Segundos:", seg+minutosASegundos(min))
}

func minutosASegundos(min minutos) segundos {
	return segundos(min * 60)
}

//sobre escribo los metodos String para cada tipo creado

func (s segundos) String() string {
	return fmt.Sprintf("%d segundos", s)
}

func (m minutos) String() string {
	return fmt.Sprintf("%d minutos", m)
}
