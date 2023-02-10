package main

import (
	"fmt"
)

func main() {
	/*for i := 1; i <= 10000; i++ {
		fmt.Println(i)
	}*/

	/*for i := 65; i <= 90; i++ {
		fmt.Println(i - 65)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}
	}*/

	/*year := 1997
	for year <= 2023 {
		fmt.Printf("%d\n", year)
		year++
	}*/

	/*year := 1997
	for {
		if year > 2023 {
			break
		}
		fmt.Println(year)
		year++
	}*/

	/*for i := 10; i < 101; i++ {
		fmt.Println(i % 4)
	}*/

	/*switch {
	case true:
		fmt.Println(true)
	case false:
		fmt.Println(false)
	}*/

	fmt.Println(true && true)
	fmt.Println(true || true)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
