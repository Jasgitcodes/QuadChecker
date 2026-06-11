package main

import "strings"

func QuadD(x, y int) string {
	if x <= 0 || y <= 0 {
		return ""
	}
	var res strings.Builder
	for row := 1; row <= y; row++ {
		for col := 1; col <= x; col++ {
			if (row == 1 && col == 1) || (row == y && col == 1) {
				res.WriteRune('A')
			} else if (row == 1 && col == x) || (row == y && col == x) {
				res.WriteRune('C')
			} else if row == 1 || row == y || col == 1 || col == x {
				res.WriteRune('B')
			} else {
				res.WriteRune(' ')
			}
		}
		res.WriteRune('\n')
	}
	return res.String()
}
