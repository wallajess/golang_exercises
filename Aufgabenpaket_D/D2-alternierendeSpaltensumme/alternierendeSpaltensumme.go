package main

// alternierendeSpaltensumme überprüft arr daraufhin, ob in der letzten
// Zeile die alternierenden Spaltensummen stehen, korrigiert diese ggf.
// und gibt die Anzahl der fehlerhaften Summen zurück.
func alternierendeSpaltensumme(arr *[anzZeilen][anzSpalten]int) int {
	// Bei geraden Reihen addieren, bei ungeraden subtrahieren
	//iterate over the rows then columns

	numWrong := 0
	var sum [anzSpalten]int

	for row := 0; row < anzZeilen-1; row++ {
		for col := 0; col < anzSpalten; col++ {
			if row%2 != 0 && row != anzSpalten-1 {
				sum[col] -= arr[row][col]
			} else {
				sum[col] += arr[row][col]
			}
		}
	}

	for i := 0; i < (anzSpalten); i++ {
		if arr[anzZeilen-1][i] != sum[i] {
			arr[anzZeilen-1][i] = sum[i]
			numWrong++
		}
	}

	return numWrong
}
