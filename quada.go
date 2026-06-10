package main

import (
	"github.com/01-edu/z01"
)

func QuadA(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}

	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			// Determine if we are at a corner
			if (i == 0 || i == y-1) && (j == 0 || j == x-1) {
				z01.PrintRune('+')
			} else if i == 0 || i == y-1 {
				// Horizontal border
				z01.PrintRune('-')
			} else if j == 0 || j == x-1 {
				// Vertical border
				z01.PrintRune('|')
			} else {
				// Inside
				z01.PrintRune(' ')
			}
		}
		// Print a newline at the end of each row
		z01.PrintRune('\n')
	}
}
