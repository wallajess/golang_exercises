package main

// bemerkenswert gibt die nächste natürliche Zahl größer start zurück, deren
// Quadrat die Summe der natürlichen Zahlen von 1 bis n ist (mit beliebigem n).
func bemerkenswert(start int) int {
	for n := start + 1; ; n++ {
		result := n * n
		sum := 0
		for i := 1; sum < result; i++ {
			sum += i
		} // sum of the first n natural numbers
		if result == sum {
			return n
		}
	}
}
