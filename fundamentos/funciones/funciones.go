package funciones

import "fmt"

func Tamanyo(cadena string) int {
	return len(cadena)
}

func Division(x, y float32) float32 {
	if y <= 0 {
		return x
	}
	return x / y
}

func Mayor(x, y int32) bool {
	if x > y {
		return true
	}
	return false
}

func Imprime(msg string) {
	fmt.Printf("%s\n")
}

func Operacion(x, y float64) (calc1, calc2 float64) {
	calc1 = x + y
	calc2 = x * y
	return calc2, calc1
}

func Doble_Triple(x int32) (doble, triple int32) {
	doble = x * 2
	triple = x * 3
	return doble, triple
}

func OperacionTwo(numero int) (int, bool) {
	if numero == 0 || numero < 0 {
		return 5, false
	}
	return 10, true
}

func Calculando(numero ...int) int {
	total := 0
	for _, num := range numero {
		total = total + num
		fmt.Printf("Total: %d\n", total)
	}
	return total
}
