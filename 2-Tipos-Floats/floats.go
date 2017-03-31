package main

import "fmt"

/*
1 - Datos Tipo Float
	float32     // 6 dígitos decimales de precisión
	float64     // 15 dígitos decimales de precisión
	complex64   // Numero complejo para float32
	complex128  // Numero complejo para float64
*/

func main() {

	var f32 float32
	var f64 float64
	var c64 complex64
	var c128 complex128

	fmt.Println("f32, f64, c64, c128 = ", f32, f64, c64, c128)

	c64 = complex(5, 6) //Asignando numero complejo mediante la funcion nativa complex
	fmt.Println("c64 = ", c64)

	c128 = complex(5, 6)
	fmt.Println("c128 = ", c128)

	fmt.Println("real - c128 * 85458.65 = ", real(c128*85458.65))       // Obtener valor de tipo real (float) con la funcion nativa real()
	fmt.Println("Imaginaria - c128 * 85458.65 = ", imag(c128*85458.65)) // Obtener valor de tipo imaginario (complex) con la funcion imag()
}
