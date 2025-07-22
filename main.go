package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	println("Pick a program")

	chooseProgram, err := reader.ReadString('\n')
	fmt.Println("Collected: ", chooseProgram)
	if err != nil {
		println(err)
		return
	}

	chooseProgram = strings.TrimSpace(chooseProgram)

	if chooseProgram == "stringers" {
		stringers()
		println("stringers")
		return
	}

	if chooseProgram == "maps" {
		maps()
		println("maps")
		return
	}

	if chooseProgram == "fibonacci closures" {
		fibonacci_closures()
		println("fibonacci_closures")
		return
	}
}
