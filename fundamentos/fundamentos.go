package main

import "fmt"

func main() {
	testA()
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
	fmt.Printf("Dato: %d \n", dato)
	fmt.Printf("Entero %d \n", entero)
	fmt.Printf("Caracter %c \n", caracter)
	fmt.Printf("Double %f \n", real)
	fmt.Printf("Flotante %f \n", flotante)
	fmt.Printf("Flotante %f \n", numero_flotante)
	fmt.Printf("Booleano %t \n", verdadero)
	fmt.Printf("Constante %d \n", TAM)
	fmt.Println("Cadena ", nulo)
	fmt.Println("Booleanos ", x, y, z)
}
