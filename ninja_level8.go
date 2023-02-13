package main

import (
	"fmt"
	"sort"
)

type user struct {
	Name string
	Age  int
}

type ByName []user

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].Name < a[j].Name }

type ByAge []user

func (b ByAge) Len() int           { return len(b) }
func (b ByAge) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b ByAge) Less(i, j int) bool { return b[i].Age < b[j].Age }

func main() {
	// #1
	// u1 := user{
	// 	Name: "Jack",
	// 	Age:  50,
	// }
	// u2 := user{
	// 	Name: "John",
	// 	Age:  35,
	// }
	// u3 := user{
	// 	Name: "Jonathan",
	// 	Age:  35,
	// }
	// users := []user{u1, u2, u3}

	// fmt.Println("Users: ", users)
	// bs, err := json.Marshal(users)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(string(bs))

	// #2
	// js := `[{"Name":"Jack","Age":50},{"Name":"John","Age":35},{"Name":"Jonathan","Age":35}]`
	// fmt.Println("JSON: ", js)
	// var users []user
	// err := json.Unmarshal([]byte(js), &users)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(users)

	// #3
	// u1 := user{
	// 	Name: "Jack",
	// 	Age:  50,
	// }
	// u2 := user{
	// 	Name: "John",
	// 	Age:  35,
	// }
	// u3 := user{
	// 	Name: "Jonathan",
	// 	Age:  35,
	// }
	// users := []user{u1, u2, u3}
	// fmt.Println("Users:  ", users)

	// err := json.NewEncoder(os.Stdout).Encode(users)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// #3
	// i := []int{3, 4, 5, 0, 1, 2}
	// str := []string{"World!", "Hello"}

	// sort.Ints(i)
	// fmt.Println(i)

	// sort.Strings(str)
	// fmt.Println(str)

	// #4
	u1 := user{
		Name: "Jack",
		Age:  50,
	}
	u2 := user{
		Name: "John",
		Age:  35,
	}
	u3 := user{
		Name: "Jonathan",
		Age:  35,
	}
	users := []user{u1, u2, u3}
	for _, val := range users {
		fmt.Println(val.Name, " : ", val.Age)
	}
	fmt.Println("\n-------Sort By Age---------")
	sort.Sort(ByAge(users))
	for _, val := range users {
		fmt.Println(val.Name, " : ", val.Age)
	}
	fmt.Println("\n-------Sort By Name--------")
	sort.Sort(ByName(users))
	for _, val := range users {
		fmt.Println(val.Name, " : ", val.Age)
	}
}
