package main

import "fmt"

const (
	a     = 42
	b int = 20
)

const (
	y1 = 2020 + iota
	y2 = 2020 + iota
	y3 = 2020 + iota
	y4 = 2020 + iota
)

func main() {
	/*var x int = 1024
	fmt.Printf("%d\n%b\n%#x\n", x, x, x)*/

	/*a := 10 == 10
	b := 10 >= 9
	c := 10 <= 9
	d := 10 > 9
	e := 10 < 9
	f := 10 != 9
	fmt.Print(a, b, c, d, e, f)*/

	// fmt.Println(a, b)

	/*i := 10
	fmt.Printf("%d\t%b\t%#X\n", i, i, i)
	j := i << 1
	fmt.Printf("%d\t%b\t%#X\n", j, j, j)*/

	/*x := `Say
	\n"hello\n" to
	 Go!`
	fmt.Print(x)*/

	fmt.Println(y1, y2, y3, y4)
}
