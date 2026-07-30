package main

// potenziere gibt den Potenzwert basis hoch exponent zurück.
// exponent muss >= 0 sein.
func potenziere(basis, exponent int) int {
	result := 1 // basis hoch 0 ist 1
	// result so oft mit der basis multiplizieren wie exponent angibt.
	for i := 1; i <= exponent; i++ {
		result = result * basis
	}
	return result
}

//Question: if exponent < i, does that mean the for-loop is not carried out?
