package main

import (
	"fmt"
	"math"
)

func main() {
	var l float64
	var w float64
	var h float64

	fmt.Println("Bitte geben Sie die drei Seitenlängen des Quaders ein:")

	_, err := fmt.Scan(&l, &w, &h)
	if err != nil {
		println("Error:", err)
		return
	}

	var vol = l * w * h
	var kant = 4 * (l + w + h)
	var obfl = 2 * (l*w + l*h + w*h)
	var umkug = 0.5 * math.Sqrt(math.Pow(l, 2)+math.Pow(h, 2)+math.Pow(w, 2))
	var raumdiag = math.Sqrt(math.Pow(l, 2) + math.Pow(h, 2) + math.Pow(w, 2))

	//add the format as a variable, so format := "%.2f"
	fmt.Printf("Ein %.2f x %.2f x %.2f Quader hat die geometrische Größen: \n", l, w, h)
	fmt.Printf("Volumen: %f \n", vol)
	fmt.Printf("Kantensumme: %f \n", kant)
	fmt.Printf("Oberfläche: %f \n", obfl)
	fmt.Printf("Umkugelradius: %f \n", umkug)
	fmt.Printf("Raumdiagonale: %f \n", raumdiag)

}
