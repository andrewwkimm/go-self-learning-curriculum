// Predict the exact output of this program. Write your answer before running it.

package main

import "fmt"

func count() {
	i := 0
	defer fmt.Println("deferred i:", i)
	i = 10
	fmt.Println("live i:", i)
}

func switchDemo(x int) {
	switch {
	case x > 0:
		fmt.Println("positive")
	case x > -5:
		fmt.Println("close to zero")
	default:
		fmt.Println("very negative")
	}
}

func main() {
	count()
	switchDemo(3)
	switchDemo(-2)
	switchDemo(-10)
}

// Two things to think through before writing your answer:

// - In count: i is 0 when defer is called, then set to 10 afterward. What does the deferred call print?
// - In switchDemo(-2): -2 > 0 is false, but -2 > -5 is true. Which case fires?

// Write your predicted output, then run it.
