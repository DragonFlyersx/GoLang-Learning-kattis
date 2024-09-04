package main

import (
	"fmt"
)

func main34() {
	var N, bag, count int
	count = 1
	fmt.Scan(&N, &bag)
	for i := 0; i < N; i++ {
		var num int
		fmt.Scan(&num)
		if bag == num {
			if count == 1 {
				fmt.Print("fyrst")
			} else if count == 2 {
				fmt.Print("naestfyrst")
			} else {
				fmt.Print(count, " fyrst")
			}
		}
		count++
	}
}
