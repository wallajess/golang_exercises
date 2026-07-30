package main

// encrypt gibt die verschlüsselte Version von klartext zurück.
// Verschlüsselung mit dem Scytale-Verfahren, Schlüssel ist das Argument scytale,
// der Durchmesser des Stabes in Buchstaben.
// scytale muss eine natürliche Zahl sein.

func encrypt(klartext string, scytale int) string {
	//determine number of rows
	var encryptedWord string
	var numCols = len(klartext) / scytale

	//Create the stab grid
	var stab [][]string
	var row []string

	//put the text into a slice of slices
	//iterate over word
	for i, letter := range klartext {
		//add each letter to an array(row)
		row = append(row, string(letter))
		//check whether the array has the prescribed length (=numCols) or if we've reached the end of klartext
		if len(row) == numCols || i == len(klartext)-1 {
			// check whether the array is filled up and if not fill it with empty strings (IS THIS NECESSARY?)
			stab = append(stab, row)
			row = []string{}
		}
	}

	//transpose it

	for col := 0; col < numCols; col++ { //iterate over the rows in each column
		for row := 0; row < len(stab); row++ { //iterate over the letters in each row
			if col < len(stab[row]) { //make sure we don't go out of bounds in rows that are shorter
				encryptedWord += string(stab[row][col]) //
			}
		}
	}
	return encryptedWord
}

// decrypt gibt die entschlüsselte Version von geheimtext zurück.
// Verschlüsselung mit dem Scytale-Verfahren, Schlüssel ist das Argument scytale,
// der Durchmesser des Stabes in Buchstaben.
// scytale muss eine natürliche Zahl sein.
func decrypt(geheimtext string, scytale int) string {

	var encryptedWord [][]string
	var decryptedWord string
	var row []string

	//iterate over the text and add each letter to a row - this time scytale is the number of columns!
	for i, letter := range geheimtext {
		row = append(row, string(letter))
		if len(row) == scytale || i == len(geheimtext)-1 {
			encryptedWord = append(encryptedWord, row)
			row = []string{}
		}
	}
	//Test to see if the matrix looks right
	//for _, row := range encryptedWord {
	//	fmt.Printf("%v\n", row)
	//}

	//

	for row := 0; row < scytale; row++ {
		for col := 0; col < len(encryptedWord); col++ {
			if row < len(encryptedWord[col]) { //
				decryptedWord += string(encryptedWord[col][row])
			}
		}
	}

	return decryptedWord
}
