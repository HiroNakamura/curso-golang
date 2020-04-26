package funciones

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
