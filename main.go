package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// Step 1: Read from stdin
	reader := bufio.NewReader(os.Stdin)
	input, _ := io.ReadAll(reader)
	content := string(input)

	if strings.TrimSpace(content) == "" {
		fmt.Println("Not a quad function")
		return
	}

	lines := strings.Split(strings.TrimRight(content, "\n"), "\n")
	y := len(lines)
	x := len([]rune(lines[0]))

	matches := []string{}

	// Step 2: Compare input to each quad
	if content == QuadA(x, y) {
		matches = append(matches, fmt.Sprintf("[quadA] [%d] [%d]", x, y))
	}
	if content == QuadB(x, y) {
		matches = append(matches, fmt.Sprintf("[quadB] [%d] [%d]", x, y))
	}
	if content == QuadC(x, y) {
		matches = append(matches, fmt.Sprintf("[quadC] [%d] [%d]", x, y))
	}
	if content == QuadD(x, y) {
		matches = append(matches, fmt.Sprintf("[quadD] [%d] [%d]", x, y))
	}
	if content == QuadE(x, y) {
		matches = append(matches, fmt.Sprintf("[quadE] [%d] [%d]", x, y))
	}

	// Step 3: Print result
	if len(matches) == 0 {
		fmt.Println("Not a quad function")
	} else {
		fmt.Println(strings.Join(matches, " || "))
	}
}
