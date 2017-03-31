package main

import "fmt"
import "strconv"

/*
1 - Datos Tipo String
	string
*/

func main() {
	cadena := "Hola Andres"
	numCaracteres := len(cadena)
	fmt.Println("Cadena:", cadena)
	fmt.Println("Numero De Caracteres:", numCaracteres)

	//obtener caracter especifico
	fmt.Println("Segundo Caracter (codigo Unicode):", cadena[1])

	//obtener rango de caracteres [inicio,cuantos espacios desde el 1]
	fmt.Println("Palabra Hola:", cadena[0:4])
	//como inicia desde 0 se puede omitir el inicio
	fmt.Println("Palabra Hola 2:", cadena[:4])
	//si se omite los dos rangos, se tomara por defecto el inicio y el final de la cadena
	fmt.Println("Omitiendo rangos:", cadena[:])

	//que se los pueda trabajar como arrays, no significa que se los pueda modificar, lo siguiente seria incorrecto:
	//cadena[0] = 'M'

	//concatenacion
	cadena += " Latorre" //=> cadena = cadena + " Latorre"
	fmt.Println(cadena)

	//crear cadena tal y como se la escribe ``
	cadena = `
	<html>
		<head>
		</head>
		<body>
			<p>hola</p>
		</body>
	</html>
	`
	fmt.Println(cadena)

	//escapar caracteres
	cadena = "numero: \"23\""
	fmt.Println(cadena)

	//convertir numero (entero) a string => strconv.Itoa()
	numero := 23
	cadena = "numero: " + strconv.Itoa(numero)

	fmt.Println(cadena)

}
