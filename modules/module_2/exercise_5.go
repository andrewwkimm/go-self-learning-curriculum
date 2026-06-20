// This program has 4 distinct errors.
// Identify every error and explain why it is wrong before making any fixes.
// Then fix all of them.
//
// Five-stage protocol:
// 1. Attempt from memory, no references, >= 10 min genuine struggle
// 2. Paste attempt here
// 3. Feedback on correctness + Uber style violations
// 4. Rewrite from scratch without referencing feedback
// 5. Add one Anki card capturing the core insight

package main

import "fmt"

// func describe(x int) {
// 	switch x { // this sets x to be compared against but x is used again in the cases; we need to get rid of this because these are independent boolean operations
// 	case x > 0:
// 		fmt.Println("positive")
// 		fallthrough // an issue because if x is positive, it executes the next case unconditionally, ie it will also print negative
// 	case x < 0:
// 		fmt.Println("negative")
// 	default:
// 		fmt.Println("zero")
// 	}
// }

// func loopDemo() {
// 	for i := 0; i <= 5; i++ {
// 		if i = 3 { // needs == and i = 3 in general is wrong, needs :=
// 			fmt.Println("three")
// 		}
// 	}
// }

// Fixed versions

func describe(x int) {
	switch {
	case x > 0:
		fmt.Println("positive")
	case x < 0:
		fmt.Println("negative")
	default:
		fmt.Println("zero")
	}
}

func loopDemo() {
	for i := 0; i <= 5; i++ {
		if i == 3 {
			fmt.Println("three")
		}
	}
}

func main() {
	describe(1)
	loopDemo()
}
