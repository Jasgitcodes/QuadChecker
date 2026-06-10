package main

import "github.com/01-edu/z01"

func QuadB(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}

	for row := 0; row < y; row++ {
		for column := 0; column < x; column++ {
			if row == 0 && column == 0 {
				z01.PrintRune('/')
			} else if row == 0 && column == x-1 {
				z01.PrintRune('\\')
			} else if row == y-1 && column == 0 {
				z01.PrintRune('\\')
			} else if row == y-1 && column == x-1 {
				z01.PrintRune('/')
			} else if row == 0 || row == y-1 || column == 0 || column == x-1 {
				z01.PrintRune('*')
			} else {
				z01.PrintRune(' ')
			}
		}

		z01.PrintRune('\n')
	}
}
