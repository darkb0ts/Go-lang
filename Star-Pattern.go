package main

import "fmt"

func main() {
	// Declare the number of rows to print
	rows := 5

	// Iterate over the rows
	for i := 1; i <= rows; i++ {

		// Iterate over the columns
		for j := 1; j <= rows-i; j++ {

			// Print a space
			fmt.Print(" ")
		}

		// Iterate over the columns
		for k := 1; k <= 2*i-1; k++ {

			// Print a star
			fmt.Print("*")
		}

		// Print a newline
		fmt.Println()
	}
}
