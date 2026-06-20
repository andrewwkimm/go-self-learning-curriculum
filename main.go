package main

import "fmt"

func classify(n int) string {
	if r := n % 2; r == 0 { // sets r to n modulo 2; if r is 0, ie no remainder, return even
		return "even"
	} else {
		_ = r // _ is blank; but why do we need this if r is always used?
		return "odd"
	}
}

func countdown(from int) {
	defer fmt.Println("liftoff")
	for i := from; i > 0; i-- {
		defer fmt.Println(i) // for loop that goes from "from" var and down by 1 until 1. since the last deferred goes first, it from from "from" to 1, then liftoff
	}
	fmt.Println("ignition") // this is printed first
}

func grade(score int) string {
	switch {
	case score >= 90:
		return "A"
	case score >= 80:
		return "B"
	case score >= 70:
		return "C"
	default:
		return "F"
	}
}

func main() {
	fmt.Println(classify(4)) // even
	fmt.Println(classify(7)) // odd
	countdown(3)             // ignition 3 2 1 liftoff
	fmt.Println(grade(85))   // B
	fmt.Println(grade(60))   // F
}
