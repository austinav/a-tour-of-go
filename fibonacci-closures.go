package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	first := false
	second := false
	lastFib := 0
	currentFib := 1
	return func() int {
		if !first {
			first = true
			return 0
		}
		if !second {
			second = true
			return 1
		}

		result := lastFib + currentFib

		lastFib = currentFib
		currentFib = result

		return result
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
