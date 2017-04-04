package main

import (
	"fmt"
)

/*
panic -> detiene la ejecucion del programa
	//crear error
	panic("descripcion")

	Ej:
	panic("Error Creado...")

recover -> recupera errores presentados en las instrucciones ejecutadas
	//recuperar Errores
	recover()

	Ej:
	//Funcion Anonima autoejecutable y ejecutandola al final
	defer func() {
		error := recover()
		fmt.Println("Errores Producidos:", error)
	}()

*/

func main() {
	fmt.Println("Inicio del programa...")
	//recuperar error o exception producido, con recover
	defer func() {
		error := recover()
		fmt.Println("Errores Producidos", error)
	}()

	panic("Error Creado manualmente") // se crea un error (exception)

}
