package main

import "fmt"

//////////////////////// #1

// func main() {
// 	a := foo()
// 	b, c := bar()
// 	fmt.Println("Result of foo: ", a)
// 	fmt.Println("Result of bar: ", b, c)
// }

// func foo() int {
// 	return 1
// }

// func bar() (string, int) {
// 	return "Hai", 2
// }

//////////////////////// #2

// func main() {
// 	i := []int{1, 2, 3, 4, 5}
// 	fmt.Println(foo(i...))
// 	fmt.Println(bar(i))
// }

// func foo(x ...int) int {
// 	total := 0
// 	for _, v := range x {
// 		total += v
// 	}
// 	return total
// }

// func bar(x []int) int {
// 	total := 0
// 	for _, v := range x {
// 		total += v
// 	}
// 	return total
// }

//////////////////////// #3

// func main() {
// 	i := []int{1, 2, 3, 4, 5}
// 	defer fmt.Println("Calling function foo: ", foo(i...))
// 	fmt.Println("Calling function bar: ", bar(i))
// }

// func foo(x ...int) int {
// 	total := 0
// 	for _, v := range x {
// 		total += v
// 	}
// 	return total
// }

// func bar(x []int) int {
// 	total := 0
// 	for _, v := range x {
// 		total += v
// 	}
// 	return total
// }

//////////////////////// #4

// type Person struct {
// 	first string
// 	last  string
// 	age   int
// }

// func (p Person) speak() {
// 	fmt.Println("Name: ", p.first, p.last, "\nAge: ", p.age)
// }

// func main() {
// 	p1 := Person{
// 		first: "Vineeth",
// 		last:  "Kumar",
// 		age:   25,
// 	}
// 	p2 := Person{
// 		first: "Anandh",
// 		last:  "Kishore",
// 		age:   30,
// 	}
// 	p1.speak()
// 	p2.speak()
// }

//////////////////////// #5

// type square struct {
// 	side float64
// }

// type circle struct {
// 	radius float64
// }

// func (s square) area() float64 {
// 	return math.Pow(s.side, 2)
// }

// const pi = 3.14

// func (c circle) area() float64 {
// 	return pi * math.Pow(c.radius, 2)
// }

// type shape interface {
// 	area() float64
// }

// func info(s shape) {
// 	switch s.(type) {
// 	case square:
// 		fmt.Println("Area of square with side ", s.(square).side, ": ", s.area())
// 	case circle:
// 		fmt.Println("Area of circle with radius ", s.(circle).radius, ": ", s.area())
// 	}
// }
// func main() {
// 	s1 := square{
// 		side: 10,
// 	}
// 	c1 := circle{
// 		radius: 10,
// 	}
// 	info(s1)
// 	info(c1)
// }

//////////////////////// #6

// func main() {
// 	func() {
// 		fmt.Println("Hai, Hello!!")
// 	}()
// }

//////////////////////// #7

// func main() {
// 	f := func() {
// 		fmt.Println("Hai, Hello!!")
// 	}
// 	f()
// }

//////////////////////// #8

// func main() {
// 	f := first()
// 	f()
// }

// func first() func() {
// 	fmt.Println("Here I am the first!!")
// 	return second
// }

// func second() {
// 	fmt.Println("Here I am the second!!")
// }

//////////////////////// #9

// func main() {
// 	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
// 	fmt.Println(sum(x...))
// 	fmt.Println(even(sum, x...))
// }

// func sum(xi ...int) int {
// 	total := 0
// 	for _, v := range xi {
// 		total += v
// 	}
// 	return total
// }

// func even(f func(xi ...int) int, yi ...int) int {
// 	y := []int{}
// 	for _, v := range yi {
// 		if v%2 == 0 {
// 			y = append(y, v)
// 		}
// 	}
// 	return f(y...)
// }

//////////////////////// #10

func main() {
	var y int
	{
		fmt.Println(y) //Not Error
		var x int      //Error
	}
	fmt.Println(x) //Error
}
