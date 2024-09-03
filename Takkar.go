package main

import "fmt"

func main15() {
	var trump, kim int
	fmt.Scan(&trump, &kim)
	if trump == kim {
		fmt.Print("WORLD WAR 3!")
	} else if trump < kim {
		fmt.Print("FAKE NEWS!")
	} else {
		fmt.Print("MAGA!")
	}

}
