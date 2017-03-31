package main

import (
	"fmt"
)

/*

ARRAYS

1 - 1 dimencion
	var nombre [tamaño]tipo
	Ej: var arr [3]int

2- multidimencion
	var nombre [tamaño][tamaño]tipo
	Ej: var arr[2][3]int
	-> [0 0 0]
	-> [0 0 0]

3 - inicializar array con tamaño definido
	var arr2 = [3]int{1, 2, 3}
	ó
	arr2 := [3]int{1, 2, 3}

4 - inicializar array con tamaño indefinido
	var arr2 = [...]int{1, 2, 3, 4}
	ó
	arr2 := [...]int{1, 2, 3, 4}

5 - asignacion
	nombre[indice] = valor

6 - tamaño de array
	len(array)

7 - comparar arrays
	a := [3]int{1,2,3}
	b := [...]int{1,2,3}
	c := [...]int{1,2,4}
	d := [4]int{1,2,3}

	a == b //=>true
	a == c //=>false
	a == d //=> error diferente tamaño (no compila)


*/

func main() {
	//filasxcolumnas
	var arr [2][3]string
	//fila 1
	arr[0][0] = "Cod"
	arr[0][1] = "Nombre"
	arr[0][2] = "Genero"

	//fila 2
	arr[1][0] = "1"
	arr[1][1] = "Andres"
	arr[1][2] = "M"

	fmt.Println(arr)

	var arr2 = [...]int{
		1,
		2,
		3}
	fmt.Println(arr2)

	a := [3]int{1, 2, 3}
	b := [...]int{1, 2, 3}
	c := [...]int{1, 2, 4}

	fmt.Println(a == b)
	fmt.Println(a == c)
}
