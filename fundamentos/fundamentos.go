package main

import (
	"fmt"
)

func main() {
	testA()
	testB()
}

func testB() {

}

func testA() {
	fmt.Println("Fundamentos de programacion en Golang")
	var dato byte = 1
	entero := int(321.0)
	caracter := '\u0000'
	var real float64 = 90.32
	var flotante float32 = 67.43
	var verdadero bool = 34 > 320
	const TAM = 1000
	numero_flotante := float64(123.8)
	var nulo string = ""
	var x, y, z bool // false false false
	//Declaracion multiple
	var (
		nombre, apellido string = "Juan", "Perez"
		exito, fracaso   bool   = true, false
		letras           string = `Alef
		Bet
		Guimmel
		Dalet
		He`
		numx, numy uint32 = 34, 54
	)

	fmt.Printf("Dato %d \n", dato)
	fmt.Printf("Entero %d \n", entero)
	fmt.Printf("Caracter %c \n", caracter)
	fmt.Printf("Double %f \n", real)
	fmt.Printf("Flotante %f \n", flotante)
	fmt.Printf("Flotante %f \n", numero_flotante)
	fmt.Printf("Booleano %t \n", verdadero)
	fmt.Printf("Constante %d \n", TAM)
	fmt.Println("Cadena ", nulo)
	fmt.Println("Booleanos ", x, y, z)
	fmt.Printf("%s %s\n", nombre, apellido)
	fmt.Printf("%t %t\n", exito, fracaso)
	fmt.Printf("Letras %s\n", letras)
	fmt.Printf("Enteros %d %d \n", numx, numy)
}
