package main

import (
	"fmt"
)

/*

1 - Operadores Aritmeticos
	+ Suma
	- Resta
	* Multiplicacion
	/ Division
	% Resto

2 - Operadores de Comparación
	== Igual que
	!= Diferente que
	< Menor que
	<= Menor o Igual que
	> Mayor que
	>= Mayor o Igual que

3 - Operadores de asignación
	     EQUIVALENCIAS
	+= | a += b | a = a + b
	-= | a -= b | a = a - b
	*= | a *= b | a = a * b
	/= | a /= b | a = a / b
	%= | a %= b | a = a % b

4 - Operadores Logicos
	&& AND
	|| OR
	! Negacion

5 - Jerarquía de Operadores
	() -> MAS IMPORTANTE
	* / %
	+ -
	== != < <= > >=
	&&
	|| -> MENOS IMPORTANTE

6 - Operadores de incremendo y decremento
	++ || a++ || a = a + 1
	-- || a-- || a = a - 1

*/

func main() {

	//Operadores de incremendo y decremento
	a := 5
	a++ // no se la puede declarar en expreciones
	fmt.Println("a++ ", a)
}
