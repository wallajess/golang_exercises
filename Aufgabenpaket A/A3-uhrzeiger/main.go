package main

import "fmt"

// Hilfsfunktionen
func winkel_hr(hr12 int, mint int) float64 {
	var winkel_hr = float64(hr12)*360/12 + float64(mint)*360/12/60
	return winkel_hr
}

func winkel_mint(mint int) float64 {
	var winkel_mint = 360 / 60 * float64(mint)
	return winkel_mint
}

func main() {
	var hr, mint int

	fmt.Println("Geben Sie bitte eine Uhrzeit an, indem Sie zunächst\n" +
		"die Stunde (von 0 bis 23) und dann die Minute (von 0 bis 59) eingeben:")

	_, errHr := fmt.Scan(&hr, &mint)
	if errHr != nil {
		fmt.Println("Fehler:", errHr)
		return
	}

	if hr < 0 || hr > 23 || mint < 0 || mint > 59 {
		fmt.Printf("%v:%v ist keine gültige Uhrzeit.\n", hr, mint)
		fmt.Println("Die Stunde muss zwischen 0 und 24 und die Minute zwischen 0 und 59 liegen.")
		return
	}

	var hr12 int
	if hr > 11 {
		hr12 = hr - 12
	} else {
		hr12 = hr
	}

	fmt.Printf("Zeigerstellung um %02d:%02d Uhr:\n", hr, mint)
	fmt.Printf("Winkel des Stundenzeigers: %.1f°\n", winkel_hr(hr12, mint))
	fmt.Printf("Winkel des Minutenzeigers:%.0f°\n", winkel_mint(mint))
}
