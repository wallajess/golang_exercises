package main

import "fmt"

func collatz(n int) {
	fmt.Printf("Die Collatz-Folge mit dem Startwert %v lautet: \n", n)
	for result := n; result != 1; {
		if result%2 == 0 {
			result /= 2
		} else {
			result = 3*result + 1
		}
		fmt.Printf("%v, ", result)
	}
	println("...\n")
}

func main() {
	collatz(19)
	collatz(23)
	collatz(42)
	collatz(122)
}
