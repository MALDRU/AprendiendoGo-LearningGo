package main //nombre del paquete cualquiera, tener en cuenta los reservados

import "fmt" // importando package fmt I/O para fmt.Printf

/*
1 - tipado
	var entero int
	var cadena string
	var boleano bool

2 - tipado automatico

	var entero = 2
	var cadena = "hola"
	var booleano = false

3 - Utilizando un solo var

	var (
		entero = 2
		cadena = "hola"
		booleano = true
	)

4 - Abreviado

	entero := 2
	cadena := "Hola"
	booleano := false

5 - mas abreviado

	entero,cadena,booleano := 2, "Hola", true

6 - asignaciones

	entero = 12
	entero,cadena = 12, "hola"

*/

func main() {
	entero, cadena, booleano := 2, "Hola", true
	fmt.Println(entero, cadena, booleano) // imprimir en nueva linea
}
