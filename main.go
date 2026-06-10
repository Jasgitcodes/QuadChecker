package main

// import "piscine"

import (
	"fmt"
	"os"
)

var (
	standardArgs []string
	x            int
	y            int
	name         string
)

func Atoi(s string) int { // convert ascii to integer
	n := 0

	for _, number := range s {
		if number < '0' || n > '9' {
			return 0
		}

		n = n*10 + int(number-'0')
	}
	return n
}

func position(s string) int { // helper function to get the length of string for looping

	return len(s)
}

func ProgramName() string { // helper function to display program name
	name := os.Args[0]

	end := position(name)
	start := 0

	for i := end - 1; i >= 0; i-- { // looping from the end of string os.args

		if name[i] == '/' || name[i] == '\\' {
			start = i + 1
			break
		}
		// programName += string(name[i])
	}

	var Name string

	for _, char := range name[start:end] { // looping from end of new name to reconvert the string
		if char == '.' {
			break
		}
		Name += string(char)
	}

	return Name
}

func main() {
	args := os.Args // recieving userinputs
	if len(args) > 1 && len(args) != 3 {
		fmt.Println("error")
		return
	}

	if len(args) == 1 {
		fmt.Println(ProgramName())
		return
	}
	// name = args[0]

	if len(args[1:]) == 2 {
		// name := ProgramName()

		x = Atoi(args[1])
		y = Atoi(args[2])
		// fmt.Printf("[%v] [%v] [%v]", name, x, y)
		QuadA(x, y)
	}

}
