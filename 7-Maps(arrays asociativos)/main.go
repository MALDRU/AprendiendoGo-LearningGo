package main

import "fmt"

/*

MAPS -> los maps son desordenados, no se imprimen como los creas, a menos que los indices sean int ordenados

1 - DECLARACION "se utiliza make para que no se inicie en nil(nulo) dandole un espacio en memoria"
	x := make(map[tipoIndice]tipovalor)
	var x = make(map[tipoIndice]tipovalor)
	o
	x := make(map[tipoIndice]tipovalor,tamaño)
	var x = make(map[tipoIndice]tipovalor,tamaño)

	//esto crearia un map nil(nulo)
	var x map[string]string

	//asi que se tendria que inicializar con valores:
	var x = map[string]string{
		"uno": "one",
		"dos": "two", // se usa coma al final o se termina en la misma linea con la llave }
	}

2 - ASIGNACION Y ACCESO
	//ASIGNACION -> para ello se debe declarar con make, o inicializar con algun valor
	x[nomIndice] = "valor"
	x[nomIndice2] = "valor2"
	Ej:
	x[1] = "andres"
	x["uno"] = "1"

	//ACCESO -> si el elemento no existe, devuelve el valor por defecto, si es int un 0 o nil si es un string
	x[nomIndice]
	Ej:
	x[1]
	x["uno"]

3 - ELIMINAR ELEMENTOS -> si el indice no existe no pasa nada, omite la operacion
	delete(map,indice)

*/

func main() {
	var mp = map[string]string{
		"uno": "one",
		"dos": "two"}

	var mp2 = map[string]int{
		"uno": 1,
		"dos": 2,
	}

	mp["tres"] = "three"
	mp["cuatro"] = "four"

	fmt.Println(mp, mp2) // lo raro es que aveces imprime en desorden :O ??

	//map devuelve 2 valores, el valor del indice y si un boolean de el resultado de la operacion

	valor, err := mp["cinco"] //en valor no asignara nada, ya que no existe y en err asignara false porque no se encontro
	fmt.Println("el resultado del elemento inexistente:")
	fmt.Println("valor:", valor)
	fmt.Println("error:", err)

	valor2, err2 := mp["tres"] //en valor asignara el correspondiente, ya que no existe y en err asignara true porque se encontro
	fmt.Println("el resultado del elemento existente:")
	fmt.Println("valor:", valor2)
	fmt.Println("error:", err2)

}
