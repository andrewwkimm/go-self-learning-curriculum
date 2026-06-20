// Write a program that prints numbers 1 through 30.
// For multiples of 3, print "Fizz" instead.
// For multiples of 5, print "Buzz" instead.
// For multiples of both, print "FizzBuzz" instead.
// Use a single for loop.
// Use an expressionless switch (no if/else chains).
// Print each result on its own line.
//
// Five-stage protocol:
// 1. Attempt from memory, no references, >= 10 min genuine struggle
// 2. Paste attempt here
// 3. Feedback on correctness + Uber style violations
// 4. Rewrite from scratch without referencing feedback
// 5. Add one Anki card capturing the core insight

package main

import "fmt"

func main() {
	for i := 1; i <= 30; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}

// func main() { // if chain version
// 	for i := 1; i <= 30; i++ {
// 		if i%3 == 0 {
// 			defer fmt.Println("Fizz")
// 		}
// 		if i%5 == 0 {
// 			defer fmt.Println("Buzz")
// 		}
// 		if i%3 == 0 && i%5 == 0 {
// 			fmt.Println("FizzBuzz")
// 		}
// 	}
// }
