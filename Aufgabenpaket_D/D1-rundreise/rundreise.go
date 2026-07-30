package main

import (
	"slices"
)

// rundreise gibt eine auf Basis der Entfernungsmatrix entfernung erstellte Route einer
// Rundreise und deren Gesamtstreckenlänge zurück.
func rundreise(entfernung [anzOrte][anzOrte]int) ([anzOrte + 1]int, int) {
	visited := make([]bool, anzOrte)
	route := []int{1}

	var distTrav int
	var entfernungS [][]int
	entfernungS = convertToSlice(entfernung)
	current := 0

	for len(route) < anzOrte {
		next, distance := nextDest(entfernungS[current], visited)

		distTrav += distance
		current = next
		visited[current] = true
		route = append(route, current+1)
	}

	distTrav += entfernungS[current][0]
	route = append(route, 1)
	OrteArray := [anzOrte + 1]int{}
	copy(OrteArray[:], route)
	return OrteArray, distTrav
}

func nextDest(currentPlace []int, visited []bool) (int, int) {
	next := -1
	visited[0] = true
	//Total distance travelled
	minDistance := slices.Max(currentPlace)
	for i, distance := range currentPlace {
		if distance != 0 && !visited[i] {
			if distance < minDistance {
				minDistance = distance
				next = i
			}
		}
	}

	// If there is only one unvisited destination left, go there
	if next == -1 {
		for i, distance := range currentPlace {
			if distance != 0 && !visited[i] {
				return i, distance
			}
		}
	}

	return next, minDistance
}

func convertToSlice(entfernung [anzOrte][anzOrte]int) [][]int {
	// Create a slice with length anzOrte
	result := make([][]int, anzOrte)

	// Copy each row from the fixed-size array into the slice
	for i := 0; i < anzOrte; i++ {
		result[i] = make([]int, anzOrte)
		copy(result[i], entfernung[i][:])
	}

	return result
}
