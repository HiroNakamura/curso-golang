package main

import (
	"fmt"
)

func main() {
	testA()
	testB()
	testC()
	testD()
}

func testD() {
	fmt.Println("***Ciclo for***")
	const MAX = 100
	for i := 0; i < MAX; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Printf("%d\n", i)
		}
	}
}

func testC() {
	fmt.Println("*** Ciclo 'while' ***")
	var (
		contador, suma int32 = 0, 0
	)

	const TAM = 100

	for contador < TAM {
		fmt.Printf("%d : %d\n", suma, contador)
		suma += contador
		contador++
		if suma >= 1000 {
			break
		}
	}
	fmt.Printf("Suma: %d\n", suma)
}

func testB() {
	var (
		exito     bool   = false
		clave     string = "12345"
		opcion    int32  = 1
		edad      int16  = 17
		operacion bool
	)

	if clave == "12345" {
		exito = true
	}
	if exito == true {
		var operacion uint32 = 30 % 10
		fmt.Println("Operacion: ", operacion)
		if operacion == 0 {
			fmt.Println("Verdadero")
		}
	}

	switch opcion {
	case 0:
		fmt.Println("1")
	case 1:
		fmt.Println("2")
	default:
		fmt.Println("3")
	}

	if edad == 18 {
		fmt.Println("Puedes votar")
	}

	if edad <= 17 {
		fmt.Println("Aun no puedes votar")
	}

	if clave == "090909" {
		fmt.Println("Son identicas")
	} else {
		fmt.Println("No son identicas")
	}

	operacion = (34 > 11) && (55 == 0) //false
	fmt.Printf("%t\n", operacion)
	operacion = (34 > 11) && (55 == 55) //true
	fmt.Printf("%t\n", operacion)
	operacion = (34 < 11) || (55 > 0) //true
	fmt.Printf("%t\n", operacion)
	operacion = !(34 < 11) && !(55 < 0) //true
	fmt.Printf("%t\n", operacion)
	operacion = !(true) //false
	fmt.Printf("%t\n", operacion)

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
