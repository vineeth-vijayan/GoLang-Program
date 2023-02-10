package main

import "fmt"

func main() {
	switch {
	case 1 == 0:
		fmt.Println("1")
	case 2 == 2:
		fmt.Println("2")
		fallthrough
	case false:
		fmt.Println("false")
		fallthrough
	case 4 == 4:
		fmt.Println("4")
		fallthrough
	default:
		fmt.Println("default")
	}
}
