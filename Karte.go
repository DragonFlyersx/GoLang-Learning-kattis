package main

import (
	"fmt"
)

func main5() {
	var s string
	fmt.Scan(&s)

	split := make(map[string]bool)
	duplicateFound := false

	for i := 0; i < len(s); i += 3 {
		segment := s[i : i+3]
		if _, exists := split[segment]; exists {
			duplicateFound = true
			break
		}
		split[segment] = true
	}
	//Kigger efter duplicate
	if duplicateFound {
		fmt.Println("GRESKA")

	} else {
		var P, K, H, T int = 13, 13, 13, 13
		for card := range split {
			switch card[0] {
			case 'P':
				P--
			case 'K':
				K--
			case 'H':
				H--
			case 'T':
				T--
			}
		}
		fmt.Println(P, K, H, T)
	}
}
