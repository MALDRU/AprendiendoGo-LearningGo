package main

import (
	"fmt"
)

/*
ESTRUCUTRAS - struct

sintaxis, segun GO
	type Nombre struct{
		Propiedad tipo
		Propiedad2 tipo
	}
	Ej:
	type Persona struct{
		Nombre string
		Genero string
	}

	Ej2: //herencia
	type Persona struct{
		Nombre string
		Genero string
	}
	type Estudiante struct{
		Persona //estudiante hereda las propiedades y metodos si los tuviera
		Carrera string
	}

Funciones a una estructura
	func (NombreEstructura) nombreFuncion(...)...{
		...
	}
	o
	func (cualquierNombre NombreEstructura) nombreFuncion(...)...{
		...
	}

	Ej:
	type Persona struct{
		Nombre,Genero string
		Peso float32
	}
	//resividor del tipo (p Persona)
	func (p *Persona) aumentarPeso(incremento float32){
		p.Peso += incremento // en struct, cuando se recibe un puntero no hay que cambiar al valor con * :O ?
	}
	...
		var pe Persona
		pe.aumentarPeso(...)
	...

Comparar Estructuras
	a == b

*/

//Persona estructura de una persona
type Persona struct {
	ID     int
	Nombre string
	Genero string
}

//Estudiante estructura de un estudiante
type Estudiante struct {
	Persona //herencia de Persona
	Carrera string
}

//Profesor estructura de un Profesor
type Profesor struct {
	Estudiante //herencia de Estudiante y Persona a la vez
	Carrera    string
}

//Automovil estructura de un Automovil
type Automovil struct {
	Nombre, Color string
}

//funcion para Automovil
func (a *Automovil) cambiarColor(color string) {
	a.Color = color //en struct, cuando se recibe un puntero no hay que cambiar al valor con * :O ?
}

func main() {
	var andres = Persona{1, "Andres", "M"}
	fmt.Println("Andres:", andres)

	angie := Persona{2, "Angie", "F"}
	fmt.Println("Angie:", angie)

	viviana := Estudiante{
		Persona{
			3,
			"Viviana",
			"F",
		},
		"Medicina",
	}

	fmt.Println("Viviana:", viviana)

	var max Estudiante
	max.ID = 4
	max.Nombre = "Max White"
	max.Genero = "M"
	max.Carrera = "Veterinario"

	fmt.Println("Max:", max)

	kakashi := Profesor{
		Estudiante{
			Persona{
				5,
				"Kakashi Hatake",
				"M",
			},
			"Hokage",
		},
		"Ninja que copia",
	}
	// kakashi := Profesor{Estudiante{Persona{5,"Kakashi Hatake","M"},"Hokage"},"Ninga que copia"} // tambien valido

	fmt.Println("Kakashi:", kakashi)

	fmt.Println("ID:", kakashi.ID)
	fmt.Println("Nombre:", kakashi.Nombre)
	fmt.Println("Carrera:", kakashi.Carrera)
	fmt.Println("Carrera Estudiante:", kakashi.Estudiante.Carrera)

	//funciones en estructuras
	nissan := Automovil{"Nissan Kics", "Blanco"}
	fmt.Println("Automovil:", nissan)

	//usar funcion de Automovil
	nissan.cambiarColor("Negro")
	fmt.Println("Automovil con cambio de color:", nissan)
}
