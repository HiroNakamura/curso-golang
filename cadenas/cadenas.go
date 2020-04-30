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
	testD()
	testE()
	testF()
	testG()
	testH()
}

func testH() {
	fmt.Println("\tSplit,ToLower, ToUpper, Title y ToTitle")
	fmt.Printf("%q\n", strings.Split("Alef,Bet, Gimmel", ","))
	titulo := strings.ToTitle("la gata del cuarto 109")
	fmt.Println(titulo)
	nombre := "FERNANDO"
	fmt.Println(strings.ToLower(nombre))
	cadena := "aldebaran"
	fmt.Println(strings.ToUpper(cadena))
	libro := "odisea en el mundo de aldebaran."
	fmt.Println(strings.Title(libro))
}

func testF() {
	fmt.Println("\tFields, Join, Repeat, Replace y ReplaceAll")
	fmt.Printf("Campos: %q\n", strings.Fields("w q r 4 5 7"))
	nombres := []string{"Ana", "Maria", "Jose", "Raul", "Georgina"}
	fmt.Println(strings.Join(nombres, "/"))
	palabra := "ba"
	fmt.Println(palabra + strings.Repeat("na", 2))
	word := strings.Replace("na na na", "k", "ky", 2)
	fmt.Println(word)
	fmt.Println(strings.ReplaceAll("La gata malbaragata", "gata", "vaca"))

}

func testG() {
	fmt.Println("\tIndex e IndexAny")
	if strings.Index("chicken", "ken") == 4 {
		fmt.Println("Correcto")
	}
	if strings.IndexAny("Chicken", "aeiouy") == 2 {
		fmt.Println("Correcto")
	}
}

func testE() {
	fmt.Println("\tEqualFold")
	if strings.EqualFold("Alfa", "alfa") == true {
		fmt.Println("Verdadero")
	}
	if strings.EqualFold("golang", "Golang") == true {
		fmt.Println("Verdadero")
	}
}

func testD() {
	fmt.Println("\tCount")
	if strings.Count("Fernando", "e") == 1 {
		fmt.Println("La letra se encuentra 1 vez")
	}
	if strings.Count("Ferrocarrilero", "r") == 5 {
		fmt.Println("La letra r se encuentra 5 veces")
	}
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
