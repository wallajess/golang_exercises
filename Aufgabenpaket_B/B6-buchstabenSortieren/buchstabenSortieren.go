package main

import "strings"

// buchstabenSortieren gibt einen String zurück, in dem alle Buchstaben
// von wort alphabetisch sortiert sind. In wort mehrfach vorkommende
// Buchstaben kommen im resultierendem String genauso oft vor. Großbuchstaben
// in wort kommen im resultierendem String als Kleinbuchstaben vor.
func buchstabenSortieren(wort string) string {

	//Put all the letters into lower case
	wordLower := strings.ToLower(wort)
	var chars []string
	for _, char := range wordLower {
		chars = append(chars, string(char))
	}

	//iterate over the word and compare two letters to see which comes first
	//first iteration takes the first letter
	for i := 0; i < len(chars); i++ {
		//second iteration for the second letter
		for j := 0; j < len(chars); j++ {
			//if the second is larger than the first, swap them around
			if chars[j] > chars[i] {
				chars[i], chars[j] = chars[j], chars[i]
			}
		}
	}
	
	// join the elements in the slice back into a string
	return strings.Join(chars, "")
}
