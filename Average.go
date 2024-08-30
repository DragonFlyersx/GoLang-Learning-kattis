package main

import "fmt"

func main8() {
	var amount_of_nubers int
	var sum = 0
	fmt.Println("Enter the amount of numbers")
	fmt.Scan(&amount_of_nubers)
	for i := 0; i < amount_of_nubers; i++ {
		var number int
		fmt.Scan(&number)
		sum = sum + number
	}
	var average = sum / amount_of_nubers
	fmt.Println("The average is:")
	fmt.Println(average)
}
