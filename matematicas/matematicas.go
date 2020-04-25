package main

import (
	"fmt"
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
}

func testA() {
	fmt.Println("Matematicas")
	var (
		x, y      int    = 0, 0
		operacion int    = 0
		opcion    string = ""
		menu      string = `
		1.Suma
		2.Resta
		3.Multiplicacion
		4.Division`
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
	default:
		fmt.Println("Ninguna operacion:")
		operacion = 0
	}
	fmt.Printf("Resultado: %d\n", operacion)
}
