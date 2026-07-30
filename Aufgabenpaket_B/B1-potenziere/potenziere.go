package main

// potenziere gibt den Potenzwert basis hoch exponent zurück.
// exponent muss >= 0 sein.
func potenziere(basis, exponent int) int {
	result := basis
	if exponent == 0 {
		return 1
	}
	for i := 1; i < exponent; i++ {
		result *= basis
	}
	return result
}

fun potenziereRekursiv(basis int, exponent int) int {
	result := basis

}
