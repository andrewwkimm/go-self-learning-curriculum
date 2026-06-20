// Modify the FizzBuzz solution from exercise 2 so that:
// 1. It reads the upper bound from os.Args[1] instead of hardcoding 30
// 2. If no argument is provided or the argument is not a valid integer,
//    print "usage: fizzbuzz <n>" and exit
// 3. Use strconv.Atoi to parse the argument
//
// Five-stage protocol:
// 1. Attempt from memory, no references, >= 10 min genuine struggle
// 2. Paste attempt here
// 3. Feedback on correctness + Uber style violations
// 4. Rewrite from scratch without referencing feedback
// 5. Add one Anki card capturing the core insight

package main

import (
	"fmt"
	"os"
	"strconv"
)

func fizzBuzz(n int) {
	for i := 1; i <= n; i++ {
		switch {
		case i%15 == 0:
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

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: fizzbuzz <n>")
		os.Exit(1)
	}

	upperBound, err := strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Println("usage: fizzbuzz <n>")
		os.Exit(1)
	}

	fizzBuzz(upperBound)
}
