package main

import (
	"fmt"
)

/*
En go el unico bucle es el For, pero este es versatil y se lo puede utilizar como todos los demas bulces conocidos

1 - ciclo FOR

Sintaxis

for variable=inicio; condicion; incremento{
	acciones por verdadero
}

2 - ciclo FOR -> para while

for condicion {
	acciones por verdadero
	accion para alterar condicion
}

3 - ciclo FOR -> para Do-while

for{
	acciones
	condicion para romper cilo -> break
}

*/

func main() {

	c := 0
	for c < 10 {
		fmt.Println("ciclo: ", c)
		c++
	}

	for c := 0; c < 10; c++ { // no hay error con la variable de arriba, por el scope o alcance, esta variable esta dentro de for, y seria otra
		fmt.Println("ciclo: ", c)
	}

}
