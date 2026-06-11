package main

import "strings"

func QuadE(x, y int) string {
	if x <= 0 || y <= 0 {
		return ""
	}
	var res strings.Builder
	for row := 1; row <= y; row++ {
		for col := 1; col <= x; col++ {
			if row == 1 && (col == 1 || col == x) {
				res.WriteRune('A')
			} else if row == y && (col == 1 || col == x) {
				res.WriteRune('C')
			} else if row == 1 || row == y {
				res.WriteRune('B')
			} else if col == 1 || col == x {
				res.WriteRune('B')
			} else {
				res.WriteRune(' ')
			}
		}
		res.WriteRune('\n')
	}
	return res.String()
}
