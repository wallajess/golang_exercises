package main

// istRechtwinklig prüft, ob das durch die Seitenlängen x, y, z gebildete
// Dreieck rechtwinklig ist.
func istRechtwinklig(x, y, z int) bool {
	return x*x+y*y == z*z || y*y+z*z == x*x || x*x+z*z == y*y
}
