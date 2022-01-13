package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func main() {

	f, _ := os.Open("day2.txt")


	valid1 := 0
	valid2 := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "-")
		min := parts[0]
		lower, err := strconv.Atoi(min)
		if err != nil {
			// insert error handling here
		  }
		moreParts := strings.Split(parts[1], " ")
		max := moreParts[0]
		higher, err := strconv.Atoi(max)
		if err != nil {
			// insert error handling here
		  }
		character := strings.Split(moreParts[1], ":")[0]
		password := moreParts[2]
		appears := strings.Count(password,character)

		if appears <= higher && appears >= lower {
			valid1 = valid1 + 1
		}
		//643

		position1 := string([]rune(password)[lower-1])
		position2 := string([]rune(password)[higher-1])
		if position1 == character && position2 != character {
			valid2 = valid2 + 1
		} else if position1 != character && position2 == character {
			valid2 = valid2 + 1
		} else {
		}

		fmt.Println(line)
		fmt.Println(min)
		fmt.Println(max)
		fmt.Println(character)
		fmt.Println(password)
		fmt.Println(position1)
		fmt.Println(position2)
	}

	fmt.Println(valid1)
	fmt.Println(valid2)



}
