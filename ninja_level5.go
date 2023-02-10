package main

import "fmt"

// type person struct {
// 	firstname             string
// 	lastname              string
// 	fav_ice_cream_flavour []string
// }

// type vehicle struct {
// 	doors  int
// 	colors string
// }

// type truck struct {
// 	vehicle
// 	fourWheel bool
// }

// type sedan struct {
// 	vehicle
// 	luxury bool
// }

func main() {
	// p1 := person{
	// 	firstname:             "Aadhi",
	// 	lastname:              "Krishnan",
	// 	fav_ice_cream_flavour: []string{"Vanilla", "Chocolate"},
	// }
	// p2 := person{
	// 	firstname:             "Kishore",
	// 	lastname:              "Raj",
	// 	fav_ice_cream_flavour: []string{"Butterscotch", "Strawberry"},
	// }
	// fmt.Println(p1.firstname)
	// fmt.Println(p1.lastname)n
	// for _, v := range p1.fav_ice_cream_flavour {
	// 	fmt.Printf("%v,", v)
	// }
	// fmt.Println()
	// fmt.Println()
	// fmt.Println(p2.firstname)
	// fmt.Println(p2.lastname)
	// for _, v := range p2.fav_ice_cream_flavour {
	// 	fmt.Printf("%v,", v)
	// }
	// fmt.Println()

	// m := map[string]person{}
	// m[p1.lastname] = p1
	// m[p2.lastname] = p2
	// for _, v := range m {
	// 	fmt.Println(v)
	// }

	// t1 := truck{
	// 	vehicle: vehicle{
	// 		doors:  2,
	// 		colors: "red",
	// 	},
	// 	fourWheel: true,
	// }
	// s1 := sedan{
	// 	vehicle: vehicle{
	// 		doors:  4,
	// 		colors: "black",
	// 	},
	// 	luxury: true,
	// }

	// fmt.Println(t1)
	// fmt.Printf("\n______\n")
	// fmt.Println(s1)

	a := struct {
		first  int
		second int
	}{
		first:  1,
		second: 2,
	}
	fmt.Println(a)
}
