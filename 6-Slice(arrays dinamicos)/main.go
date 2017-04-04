package main

import "fmt"

/*
Slice -> array que puede cambiar su tamaño en tiempo de ejecucion, siendo no necesario indicar su tamaño

1 - declaracion
	var slice []int

2 - inicializacion
	var slice = []int{1, 2, 3}

3 - longitud -> numero de elementos actuales
	len(array)

4 - capacidad -> numero de elementos que puede expadirse o maximo
	cap(array)

5 - declarando y especificando len con make
	var slice = make([]int, 3) //=> len = 3

6 - declarando y especificando len y cap
	var slice = make([]int, 3, 100) //=> len = 3 y cap = 100

7 - agregar elemento "append"
	slice = append(slice,elementos)
	Ej:
	var slice = []string
	slice = append(slice,"elemento1")
	slice = append(slice,"elemento2","elemento3","elemento4")

8 - copiar elementos "copy"
	origen := []int{1,2,3}
	destino := []int{3,4,3}
	copy(destino,origen) // 1 en donde se copiara, 2 de donde se copiara

	//si el slice en destino es mas pequeño que el origen, solo se copian los que alcancen

*/

func main() {
	var slice []string
	slice = append(slice, "elemento1")
	slice = append(slice, "elemento2", "elemento3", "elemento4")
	fmt.Println(slice)

	origen := []int{1, 2, 3}
	destino := []int{3, 4, 3}
	copy(destino, origen)
	fmt.Println(destino)
}
