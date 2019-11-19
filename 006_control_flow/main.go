package main

import "fmt"

func main() {
	door := "дверь"

	x := 1
	for ; x <= 10; x++ {

		//fmt.Println(x%2 == 0)

		if x%2 == 0 {
			fmt.Println(x)
			continue
			//break
		}

		fmt.Println("not continue")
	}

	if door != "test" {
		fmt.Println(door)
	} else if door != "ololo" {
		fmt.Println("not ololo")
	} else {
		fmt.Println("else")
	}

	switch door {
	case "ololo":
		fmt.Println(door)
	default:
		fmt.Println("default")
	}

}
