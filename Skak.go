package main

import "fmt"

func main33() {
	var Rookx, Rooky int
	var pawnx, pawny int
	fmt.Scan(&Rookx, &Rooky, &pawnx, &pawny)

	if Rookx == pawnx || Rooky == pawny {
		fmt.Println("1")
	} else {
		fmt.Println("2")
	}

}
