package main

import "fmt"

func main7() {
	var Apple float64
	var orange float64
	var Pineapple float64
	var Apple_con float64
	var orange_con float64
	var Pineapple_con float64

	fmt.Scan(&Apple, &orange, &Pineapple, &Apple_con, &orange_con, &Pineapple_con)

	// math
	var total_con = Apple_con + orange_con + Pineapple_con

	// find limiting factor

	var limiting_factor = 0.0
	var Apple_coktail = Apple / (Apple_con / total_con)
	var orange_coktail = orange / (orange_con / total_con)
	var Pineapple_coktail = Pineapple / (Pineapple_con / total_con)

	if Apple_coktail < orange_coktail {
		limiting_factor = Apple_coktail
	} else {
		limiting_factor = orange_coktail
	}
	if limiting_factor > Pineapple_coktail {
		limiting_factor = Pineapple_coktail
	}

	var apple_left = Apple - limiting_factor*(Apple_con/total_con)
	var orange_left = orange - limiting_factor*(orange_con/total_con)
	var Pineapple_left = Pineapple - limiting_factor*(Pineapple_con/total_con)
	// then Print
	fmt.Println(apple_left, orange_left, Pineapple_left)
}
