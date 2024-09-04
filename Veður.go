package main

import "fmt"

func main() {
	var wind, roads int
	fmt.Scan(&wind, &roads)
	for i := 0; i < roads; i++ {
		var roadname string
		var windspeed int
		fmt.Scan(&roadname, &windspeed)
		if windspeed < wind {
			fmt.Println(roadname + " lokud")
		} else {
			fmt.Println(roadname + " opin")
		}
	}
}
