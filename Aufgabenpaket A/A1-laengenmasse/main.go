package main

import "fmt"

func main() {
	//declare variables
	var lenM float64

	fmt.Println("Geben Sie eine Länge (in Metern) ein:")

	fmt.Scan(&lenM)

	var lenMM float64
	lenMM = lenM * 1000

	var lenKM float64
	lenKM = lenM / 1000

	var zoll float64
	zoll = lenM * 100 / 2.54

	var seemeilen float64
	seemeilen = lenM / 1852

	var lichtjahre float64
	lichtjahre = lenM / 9460730472580800

	fmt.Println(lenM, "Meter entsprechen:") //Printf means formatted string, %1.3e is the formating for 3 places behind the decimal and exponent
	fmt.Printf("%1.3e mm \n", lenMM)
	fmt.Printf("%1.3e km \n", lenKM)
	fmt.Printf("%1.3e Zoll \n", zoll)
	fmt.Printf("%1.3e sm \n", seemeilen)
	fmt.Printf("%1.3e Lj \n", lichtjahre)
}
