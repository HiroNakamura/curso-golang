package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("*** Cadenas ***")
	testA()
	testB()
	testC()
}

func testC() {
	fmt.Println("\tContainsAny")
	if strings.ContainsAny("Team", "eu") == true {
		fmt.Println("Correcto: al menos hay una e")
	}
	if strings.ContainsAny("Golang", "i") != true {
		fmt.Println("Correcto: no se encuentra alguna i")
	}
	if strings.ContainsAny("Futuro", "ou") == true {
		fmt.Println("Correcto:  se encuentra una o")
	}
}

func testB() {
	fmt.Println("\tContains")
	if strings.Contains("En un mundo celestial", "mundo") == true {
		fmt.Println("La cadena contiene la palabra: mundo")
	}
	if strings.Contains("Las huertas producen las mejores zanahorias", "huertas") == true {
		fmt.Println("La cadena contiene la palabra: huerta")
	}
	if strings.Contains("Genesis Haumeaga", "Apocalipsis") != true {
		fmt.Println("Correcto: no se encuenta la palabra Apocalipsis")
	}
}

func testA() {
	fmt.Println("\tCompare")
	if strings.Compare("A", "A") == 0 {
		fmt.Println("Verdadero A = A --> 0")
	}

	if strings.Compare("B", "X") == 1 {
		fmt.Println("Verdadero B != A --> 1")
	}

	if strings.Compare("X", "B") == -1 {
		fmt.Println("Verdadero X != B --> -1")
	}
}
