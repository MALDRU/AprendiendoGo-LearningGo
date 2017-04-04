package main

/*
Punteros -> direcciones de memoria asignadas a variables, su valor por defecto es nil

//crear puntero
	a := new(int)

//crear Puntero con & -> obtiene la direccion de la variable
	a := 23
	b := &a //direccion de a



//obteniendo valor de la direccion o puntero
	a := 23
	b := &a //direccion de a
	fmt.println(*b) //valor de la direccion

//en funciones, para recibir un puntero se lo hace con *
	func sumar(punt *int){ //recibe un puntero de int
		//aca funciona normal: con * se recupera el valor, y con & se recupera la direccion
	}

*/

func sumarDos(num *int) {
	*num += 2
}

func main() {
	a := 23
	println("Numero Inicializado: ", a) // funciona sin fmt.Print...  :O ??
	println("Puntero Del Numero Inicializado: ", &a)

	b := &a
	println("Variable apuntado a la direccion: ", b)
	println("Recuperando el valor de la direccion: ", *b)

	*b = 10 // accedo al valor de la direccion y se lo cambio
	println("Nuevo valor de a: ", a)

	sumarDos(&a)
	println("Sumar 2 mediante la funcion:", a)

}
