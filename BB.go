package main

import (
	"fmt"
	"strconv"
)

func main6() {
	var n int

	var count int
	var bad bool
	bad = false
	count = 1
	fmt.Scan(&n)
	for i := 1; n > i-1; i++ {
		var s string
		fmt.Scan(&s)

		converteds, err := strconv.Atoi(s)
		if err == nil {
		}
		if s == "mumble" { // skal compare string om det er mumble
			count++

		} else if converteds != count {
			bad = true // int skal passe med count ellers fejl

		} else {
			count++
		}

	}

	if bad == true {
		fmt.Println("something is fishy")
	} else {
		fmt.Println("makes sense")
	}
}
