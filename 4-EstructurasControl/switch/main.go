package main

import (
	"fmt"
)

/*

switch

No es necesario escribir el break para romper por cada case, porque ya lo pone por defecto

pero si se necesita evaluar todas los case o que continue se utiliza "fallthrough"
Ej:
case 1:
	acciones
	fallthrough // continua con el siguiente case si dado el caso entra en este case
case 2:
	acciones // aca no continua si entra aca, ya que si no se especifica por defecto estaria el break y sale
case 3:
	acciones

SINTAXIS

1 - NORMAL
switch valor a evalular{
	case valor:
		acciones
	case valor2:
		acciones
	default:
		acciones
}

2 - PARA VARIAS CONDICIONES
switch {
	case condicion:
		acciones
	case condicion2:
		acciones
	default:
		acciones
}

*/

func main() {
	var dia int8
	fmt.Println("Digita un numero de la semana: ")
	fmt.Scanln(&dia)

	switch dia {
	case 1:
		fmt.Println("Lunes")
	case 2:
		fmt.Println("Martes")
	case 3:
		fmt.Println("Miercoles")
	case 4:
		fmt.Println("Jueves")
	case 5:
		fmt.Println("Viernes")
	case 6:
		fmt.Println("Sabado")
	case 7:
		fmt.Println("Domingo")
	default:
		fmt.Println("Ningun dia de la semana corresponde con lo que ingresaste")
	}
}
