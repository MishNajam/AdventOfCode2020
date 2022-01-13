package main

import (
	"fmt"
	"bufio"
	"os"
	// "strings"
	// "strconv"
)

func main() {

	f, _ := os.Open("day3.txt")

	open := 0
	trees := 0
	x := 1
	// y := 1

	xcoord := 0
	// ycoord := 0

	lineCount := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		length := len([]rune(line))

		currentPosition := string([]rune(line)[xcoord])

		if lineCount % 2 == 0 {
			if currentPosition == "." {
				open = open + 1
			} else if currentPosition == "#" {
				trees = trees + 1
			}
			xcoord = (xcoord + x) % length
		}
		lineCount = lineCount + 1

		fmt.Println(line)
		fmt.Println(length)
		fmt.Println(trees)
		fmt.Println(open)
	}
	fmt.Println(trees)

	fmt.Println(60*191*64*63*32)

	//slope 1 : 60
	//slope 2 : 191
	//slope 3 : 64
	//slope 4 : 63
	//slope 5 : 38
}
