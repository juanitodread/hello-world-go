package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// create a map(string, int) map and store each new line
	counts := make(map[string]int)

	// create a new scanner object to read from stdin
	inputs := bufio.NewScanner(os.Stdin)

	// read the line as set as the key of the map, every time the key is
	// repeated increment the counter
	for inputs.Scan() {
		counts[inputs.Text()]++
	}

	// print results
	for line, count := range counts {
		if count > 1 {
			fmt.Printf("The line \"%s\" was introduced %d times\n", line, count)
		}
	}
}
