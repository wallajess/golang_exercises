package main

// istPerfekt bestimmt, ob zahl eine perfekte Zahl ist.
// zahl muss eine natürliche Zahl sein.
func istPerfekt(zahl int) bool {
	// bitte korrigieren Sie den Rumpf dieser Funktion.
	var teiler int
	for i := 1; i < zahl; i++ {
		if zahl%i == 0 {
			teiler += i
		}
	}
	return teiler == zahl
}
