package main

import "fmt"

// /////////////////////////// #1

// func main() {
// 	var a int = 10
// 	fmt.Printf("%d\n", &a)
// 	fmt.Printf("%f\n", &a)
// 	fmt.Printf("%o\n", &a)
// 	fmt.Printf("%x\n", &a)
// 	fmt.Printf("%u\n", &a)
// }

/////////////////////////// #2

type Person struct {
	first string
	last  string
	age   int
}

func changeMe(p *Person) {
	p.first = "Vineeth"
	p.last = "Kumar"
	p.age = 25
}

func main() {
	p1 := Person{
		first: "Dr",
		last:  "Strange",
		age:   150,
	}
	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)
}
