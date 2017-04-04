package main

import "fmt"

/*
FUNCIONES -> si comienza con minuscula es "No exportada" (privada), con mayuscula "Exportada" (publica) y es recomendable agregar un comentario

1 - DECLARACION

	//sin retorno, ni parametros
	func nombreFuncion () {
		//acciones
	}

	//con parametros
	func nombreFuncion(parametro1 tipo, parametro2 tipo){
		//acciones
	}
	o
	func nombreFuncion(parametro1, parametro2 tipo){ // si son del mismo tipo solo se lo pone al final
		//acciones
	}


	//con parametros y retorno
	func nombreFuncion(parametro1 tipo, parametro2 tipo) tipo{
		return //acciones
	}
	o
	func nombreFuncion(parametro1 tipo, parametro2 tipo) (nombre tipo){
		//acciones
		return //no es necesario especificar el retorno, ya que go toma el valor del nombre del retorno y lo devuelve
	}
	o
	func nombreFuncion(parametro1 tipo, parametro2 tipo) (nombre tipo, nombre2 tipo){
		//acciones
		return //no es necesario especificar el retorno, ya que go toma los valores del los retornos y lo devuelve
	}
	o
	func nombreFuncion(parametro1 tipo, parametro2 tipo) (tipo1,tipo2){
		//acciones
		return valorTipo1,valorTipo2//aca se especifica los valores, ya que solo se especifco el tipo, y estos valores deben estar en el mismo orden de la declaracion
	}

2 - ORDEN DE EJECUCION

	//ejecutar una funcion despues de que todo ya se halla ejecutado "defer"

	defer func1 (); // se ejecutara de ultima
	func2 ();

*/

func dividir(num1, num2 float32) (r float32, err string) {
	if num2 == 0 {
		err = "El divisor no puede ser 0"
	} else {
		r = num1 / num2
	}
	return
}

func cerrarTodo() {
	fmt.Println("Listo, Ya cerre Todo :)")
}

func main() {

	defer cerrarTodo() // se ejecuta cuando todo halla terminado

	r, err := dividir(6, 2)
	fmt.Println("La operacion trajo un error: ", err)
	fmt.Println("El resultado es:", r)

	r, err = dividir(6, 0)
	fmt.Println("La operacion trajo un error: ", err)
	fmt.Println("El resultado es:", r)

	//omitir un valor de retorno con _
	r, _ = dividir(29, 4)
	fmt.Println("El resultado es:", r)
}
