package main

import "fmt"

func main() {
	// x := [5]int{1, 2, 3, 4, 5}
	// for i, _ := range x {
	// 	fmt.Printf("%v\t\n", x[i])
	// }
	// fmt.Printf("%T\n", x)

	// x := []int{1, 2, 3, 4, 5}
	// for i, _ := range x {
	// 	fmt.Printf("%v\t\n", x[i])
	// }
	// fmt.Printf("%T\n", x)

	// x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// fmt.Printf("%v\n", x[:5])
	// fmt.Printf("%v\n", x[5:10])
	// fmt.Printf("%v\n", x[2:7])
	// fmt.Printf("%v\n", x[1:6])

	// x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// fmt.Println(x)
	// fmt.Println("Length: ", len(x), "\nCapacity: ", cap(x))
	// x = append(x, 11)
	// fmt.Println(x)
	// fmt.Println("Length: ", len(x), "\nCapacity: ", cap(x))
	// x = append(x, 12, 13, 14, 15)
	// fmt.Println(x)
	// fmt.Println("Length: ", len(x), "\nCapacity: ", cap(x))
	// y := []int{16, 17, 18, 19, 20}
	// x = append(x, y...)
	// fmt.Println(x)
	// fmt.Println("Length: ", len(x), "\nCapacity: ", cap(x))

	// x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// x = append(x[:2], x[6:]...)
	// fmt.Println(x)

	// x := make([]int, 5, 5)
	// fmt.Println(len(x), cap(x))
	// x = []int{1, 2, 3, 4}
	// fmt.Println(len(x), cap(x))

	// fr := []string{"James", "Bond", "Shaken, not Stirred"}
	// sr := []string{"Miss", "Moneypenny", "Helloooooo, James"}
	// ss := [][]string{fr, sr}
	// fmt.Println(ss)

	// for _, row := range ss {
	// 	for _, ele := range row {
	// 		fmt.Println(ele)
	// 	}
	// }

	// m := map[string][]string{
	// 	"bond_james":      []string{"shaken", "martines", "cake"},
	// 	"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
	// 	"no_dr":           []string{"Being Evil", "Ice Cream", "Sunsets"},
	// }
	// // m["bond_james"] = []string{"shaken", "martines", "cake"}
	// // m["moneypenny_miss"] = []string{"James Bond", "Literature", "Computer Science"}
	// // m["no_dr"] = []string{"Being Evil", "Ice Cream", "Sunsets"}
	// fmt.Println(m)
	// for k, v := range m {
	// 	fmt.Println("This is a record for ", k, ":\t")
	// 	for _, val := range v {
	// 		fmt.Printf("\t%v\n ", val)
	// 	}
	// 	fmt.Printf("\n")
	// }

	// m := map[string][]string{
	// 	"bond_james":      []string{"shaken", "martines", "cake"},
	// 	"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
	// 	"no_dr":           []string{"Being Evil", "Ice Cream", "Sunsets"},
	// }
	// // m["bond_james"] = []string{"shaken", "martines", "cake"}
	// // m["moneypenny_miss"] = []string{"James Bond", "Literature", "Computer Science"}
	// // m["no_dr"] = []string{"Being Evil", "Ice Cream", "Sunsets"}
	// m["kumar_vineeth"] = []string{"games", "sweets", "workout"}
	// fmt.Println(m)
	// for k, v := range m {
	// 	fmt.Println("This is a record for ", k, ":\t")
	// 	for _, val := range v {
	// 		fmt.Printf("\t%v\n ", val)
	// 	}
	// 	fmt.Printf("\n")
	// }

	m := map[string][]string{
		"bond_james":      []string{"shaken", "martines", "cake"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr":           []string{"Being Evil", "Ice Cream", "Sunsets"},
	}
	// m["bond_james"] = []string{"shaken", "martines", "cake"}
	// m["moneypenny_miss"] = []string{"James Bond", "Literature", "Computer Science"}
	// m["no_dr"] = []string{"Being Evil", "Ice Cream", "Sunsets"}
	m["kumar_vineeth"] = []string{"games", "sweets", "workout"}
	fmt.Println(m)
	delete(m, "bond_james")
	for k, v := range m {
		fmt.Println("This is a record for ", k, ":\t")
		for _, val := range v {
			fmt.Printf("\t%v\n ", val)
		}
		fmt.Printf("\n")
	}
}
