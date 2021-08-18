package main

import "fmt"

const PI = 3.142

type Gender int
const (
	Male Gender = iota
	Female 
	Other
	NewGender
	NewStuff
)

func main() {
	var tim Gender
	var jane Gender
	tim = NewStuff
	jane = NewGender
	fmt.Println(tim)
	fmt.Println(jane)
	fmt.Println(jane == tim)


	// fmt.Println(add(6, 8))
	// star := add(3, 4)
	// fmt.Println(star)

	// var adder AddI

	// adder = func(x, y int) int {
	// 	return x + y
	// }

	// fmt.Println(adder(7, 100))
}

func add(a, b int) int {
	return a + b
}

type AddI func(int, int) int

func test(a, n int) int {
	return a + n
}

func testII() {
	fmt.Println(test(10, 10))
}

func (g Gender) toString() string {
	switch g {
	case Male:
		return "Male"
	case Female:
		return "Female"
	case Other:
		return "Other"
	case NewGender:
		return "NewGender"
	default:
		return "Unknown Gender"
	}
}