package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("you can do nothing to stop the time")
	fmt.Println("it will fools you at the most devastating times\n")
	fmt.Println("no one told you when to run,\nyou missed the starting gun.\n")
	path, _ := os.Getwd()

	fmt.Print("current path: ", path, "\n")
}
