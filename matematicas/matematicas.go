package main

import (
	"fmt"
	"math"
)

func suma(x, y int) int {
	return x + y
}

func resta(x, y int) int {
	return x - y
}

func multiplicacion(x, y int) int {
	return x * y
}

func division(x, y int) int {
	if y == 0 || y < 0 {
		return x
	}
	return x / y
}

func main() {
	testA()
	testB()
}

func testB() {
	fmt.Println("*** Usando funciones math ***")
	const TAM = 500
	var (
		x, y, operacion float64
	)

	x, y = 3.0, -6.0
	fmt.Printf("x: %.0f, y: %.0f\n", x, y)
	operacion = math.Max(x, float64(y))
	fmt.Println("Max: ", operacion)
	operacion = math.Min(x, float64(y))
	fmt.Println("Min: ", operacion)
	operacion = math.Atan2(y, x)
	fmt.Println("Atan2: ", operacion)
	operacion = math.Exp(x)
	fmt.Println("Exp: ", operacion)
	operacion = math.Abs(y)
	fmt.Println("Abs: ", operacion)
	operacion = math.Acos(x + y)
	fmt.Println("Acos: ", operacion)
	operacion = math.Acosh(y)
	fmt.Println("Acosh: ", operacion)
	operacion = math.Ceil(x)
	fmt.Println("Ceil: ", operacion)
	operacion = math.Log(y)
	fmt.Println("Log: ", operacion)
	operacion = math.Round(y)
	fmt.Println("Round: ", operacion)
	operacion = math.Sqrt(x)
	fmt.Println("Sqrt: ", operacion)
	operacion = math.Mod(x, math.Abs(y))
	fmt.Println("Mod: ", operacion)
	operacion = math.Pow(y, x)
	fmt.Println("Pow: ", operacion)
}

func testA() {
	fmt.Println("*** Matematicas ***")
	var (
		x, y      int    = 0, 0
		operacion int    = 0
		opcion    string = ""
		menu      string = `
		1.Suma
		2.Resta
		3.Multiplicacion
		4.Division
		5.Salir
		`
	)

	fmt.Println("\tMenu")
	fmt.Printf("%s\n", menu)
	fmt.Println("Escoge una opcion:")
	fmt.Scanf("%s", &opcion)

	fmt.Println("Introduce dos numeros:")
	fmt.Scan(&x, &y)
	fmt.Printf("x: %d, y: %d\n", x, y)

	switch opcion {
	case "1":
		fmt.Println("Suma:")
		operacion = suma(x, y)
	case "2":
		fmt.Println("Resta:")
		operacion = resta(x, y)
	case "3":
		fmt.Println("Multiplicacion:")
		operacion = multiplicacion(x, y)
	case "4":
		fmt.Println("Division:")
		operacion = division(x, y)
	case "5":
		fmt.Println("Salimos")
		break
	default:
		fmt.Println("Ninguna operacion:")
		operacion = 0
	}

	fmt.Printf("Resultado: %d\n", operacion)
}
