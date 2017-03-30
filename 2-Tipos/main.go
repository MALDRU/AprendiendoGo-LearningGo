package main

import "fmt"
import "unsafe"

/*
1 - Enteros con signo
	var entero8 int8   //Todos los enteros de 8-bit con signo (-128 to 127)
	var entero16 int16 //Todos los enteros de 16-bit con signo (-32768 to 32767)
	var entero32 int32 //Todos los enteros de 32-bit con signo (-2147483648 to 2147483647)
	var entero64 int64 //Todos los enteros de 64-bit con signo (-9223372036854775808 to 9223372036854775807)


2 - Enteros sin signo
	var uentero8 uint8   //Todos los enteros de 8-bit sin signo (0 to 255)
	var uentero16 uint16 //Todos los enteros de 16-bit sin signo (0 to 65535)
	var uentero32 uint32 //Todos los enteros de 32-bit sin signo (0 to 4294967295)
	var uentero64 uint64 //Todos los enteros de 64-bit sin signo (0 to 18446744073709551615)

3 - Alias
	var enteroByte byte //alias para uint8
	var enteroRune rune //alias parar int32

4 - Tipos dependiente de la plataforma
	var enteroUint uint       //32 o 64 bits
	var enteroInt int //32 o 64 bits
	var enteroUintptr uintptr //Entero sin signo lo suficientemente grande para contener el valor de un puntero.

5 - conversiones de tipos

	entero32 = 60
	entero64 = 120

	fmt.Println(entero32 + int32(entero64))
*/

func main() {
	var entero32 int32 = 60
	var entero64 int64 = 120

	fmt.Println(entero32 + int32(entero64))

	//obtener tamaño en bytes del tipo de una variable
	//entero de 32 bits en bytes 4 (1 byte => 8 bits)
	fmt.Println(unsafe.Sizeof(entero32))
	//entero de 64 bits en bytes 8dDSAF
	fmt.Println(unsafe.Sizeof(entero64))
}
