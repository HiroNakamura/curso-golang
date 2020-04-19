package main

import "fmt"

func main() {
	testA()
}

func testA() {
	fmt.Println("Fundamentos de programacion en Golang")
	var dato byte = 1
	entero := int("321")
	caracter := '\u0000'
	var real double = 90.32
	var flotante float32 = 67.43f
	var verdadero bool = 34 > 320
	const TAM = 1000
	numero_flotante = float32("123.8")
	fmt.Printf("Dato: %d\n",dato)
	fmt.Printf("Entero %d\n",entero)
	fmt.Printf("Caracter %c\n",caracter)
	fmt.Printf("Double %d\n",real)
	fmt.Printf("Flotante %f\n",flotante)
	fmt.Printf("Flotante %\n",numero_flotante)
	fmt.Printf("Booleano %b\n",verdadero)
	fmt.Printf("Constante %d\n",TAM)
}

