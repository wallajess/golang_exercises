package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	const maxNum = 100 //Damit man das für das gesamtProgramm leichter mal ändern kann, aber nicht im Laufe des Programms ändern kann
	num := rand.Intn(maxNum + 1)
	var guess int
	counter := 1
	var previousDiff int
	var difference int

	fmt.Println("Versuchen Sie eine natürliche Zahl zu erraten. Sie liegt zwischen 0 und 100.")

	for i := 0; guess != num; i++ {
		fmt.Printf("Ihr %d. Versuch\n", counter)
		_, err := fmt.Scan(&guess)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		difference = int(math.Abs(float64(guess - num)))
		println("difference is", difference)

		if i == 0 {
			switch {
			case difference <= 5:
				println("heiß")
			case difference > 5:
				println("kalt")
			}

		} else {
			switch {
			case difference < 6:

				println("heiß")
			case difference < previousDiff:
				println("wärmer")
			case difference > previousDiff:
				println("kälter")
			case difference == previousDiff:
				println("gleicher Abstand")
			}
		}
		counter++
		previousDiff = difference
	}

	fmt.Printf("Herzlichen Glückwunsch! Sie haben die Zahl in %d Versuchen erraten.\n", counter+1)

}
