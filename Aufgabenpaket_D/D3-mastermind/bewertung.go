package main

// bewertung gibt die Bewertung von versuch an Hand von code zurück.
// Die zurückgegebene Bewertung ist ein Paar (schwarze, weiße).
func bewertung(versuch, code [anzStellen]string) (int, int) {

	var schwarz int
	var weiss int

	for i := 0; i < len(versuch); i++ {

		if versuch[i] == code[i] {
			schwarz++
		} else if versuch[i] != code[i] {
			for j := i + 1; j < len(versuch); j++ {
				if versuch[i] == code[j] {
					weiss++
				}
			}
		}
	}
	return schwarz, weiss
}
