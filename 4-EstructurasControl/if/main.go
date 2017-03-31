package main

import (
	"fmt"
)

/*
1 - IF

Sintaxis

if condicion {
	acciones por verdadero
}

if condicion {
	acciones por verdadero
}else{
	acciones por falso
}

if condicion {
	acciones por verdadero
}else if{
	acciones por verdadero2
}else
{
	acciones por falso
}

*/

func main() {

	c := 0
	if c > 0 {
		fmt.Printf("verdadero")
	} else {
		fmt.Printf("falso")
	}

}
