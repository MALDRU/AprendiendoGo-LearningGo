package main

import (
	"fmt"
)

func main() {

	fmt.Print(`
---------------------------------
|      ¿ Es Par o Impar ?       |
---------------------------------

Digita el numero a verificar: `)

	//capturar ingreso por consola
	var numero int
	fmt.Scan(&numero)

	//validando si es par o impar o neutro
	fmt.Println("\n---------------------------------")
	if numero == 0 {
		fmt.Printf("El numero %d es neutro", numero)
	} else if numero%2 == 0 {
		fmt.Printf("El numero %d es Par", numero)
	} else {
		fmt.Printf("El numero %d es Impar", numero)
	}
	fmt.Println("\n---------------------------------\n")
}
