package main

import (
	"github.com/01-edu/z01"
)

func QuadE(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}

	dx := x - 1
	dy := y - 1
	for row := 0; row < y; row++ {
		for column := 0; column < x; column++ {
			if row == 0 && column == 0 {
				z01.PrintRune('A')
			} else if row == 0 && column == dx {
				z01.PrintRune('C')
			} else if row == dy && column == 0 {
				z01.PrintRune('C')
			} else if row == dy && column == dx {
				// z01.PrintRune('A')
				z01.PrintRune('A')
			} else if row == 0 || column == 0 || row == dy || column == dx {
				// fmt.Print("B")
				z01.PrintRune('B')
			} else {
				// fmt.Print(" ")
				z01.PrintRune(' ')
			}
		}

		z01.PrintRune('\n')
	}
}
