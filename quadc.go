package main

import "github.com/01-edu/z01"

func QuadC(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}

	dy := y - 1
	dx := x - 1
	for row := 0; row < y; row++ {
		for column := 0; column < x; column++ {
			if row == 0 && column == 0 {
				z01.PrintRune('A')
			} else if row == 0 && column == dx {
				z01.PrintRune('A')
			} else if row == dy && column == 0 {
				z01.PrintRune('C')
			} else if row == dy && column == dx {
				z01.PrintRune('C')
			} else if row == 0 || row == dy || column == 0 || column == dx {
				z01.PrintRune('B')
			} else {
				z01.PrintRune(' ')
			}
		}

		z01.PrintRune('\n')
	}
}
