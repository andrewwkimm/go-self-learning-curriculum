// Write a function runWithCleanup(name string) that:
// 1. Prints "starting: <name>"
// 2. Defers a print of "done: <name>"
// 3. If name is empty, prints "error: name required" and returns early
// 4. Otherwise prints "working: <name>"
//
// Call it twice from main: once with "task-1", once with "".
// The defer must fire on both code paths.
//
// Five-stage protocol:
// 1. Attempt from memory, no references, >= 10 min genuine struggle
// 2. Paste attempt here
// 3. Feedback on correctness + Uber style violations
// 4. Rewrite from scratch without referencing feedback
// 5. Add one Anki card capturing the core insight

package main

import "fmt"

func runWithCleanup(name string) {
	defer fmt.Println("done:", name)

	fmt.Println("starting:", name)

	if name == "" {
		fmt.Println("error: name required")
		return
	}

	fmt.Println("working:", name)
}

func main() {
	runWithCleanup("task-1")
	runWithCleanup("")
}
