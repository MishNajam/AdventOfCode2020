package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func contains(s []int, str int) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func ReadLine(lineNum int) (line string) {
	f, _ := os.Open("day8.txt")
	sc := bufio.NewScanner(f)
	counter := 0
    for sc.Scan() {
        if counter == lineNum {
            return sc.Text()
		}
		counter = counter + 1
    }
    return line
}

func main() {

	f, _ := os.Open("day8.txt")

	scanner := bufio.NewScanner(f)
	jmpCount := 0
	nopCount := 0
	docLength := 0
	listOfJmp := []int{}
	listOfNop := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		operation := parts[0]
		if operation == "jmp" {
			jmpCount = jmpCount + 1
			listOfJmp = append(listOfJmp, docLength)
		} else if operation == "nop" {
			nopCount = nopCount + 1
			listOfNop = append(listOfNop, docLength)
		}
		docLength = docLength + 1
	}

	fmt.Println(docLength)
	fmt.Println(jmpCount)
	fmt.Println(listOfJmp)
	fmt.Println(nopCount)
	fmt.Println(listOfNop)



	for i := range listOfJmp {
		accumulator := 0
		lineNumber := 0
		listOfLines := []int{}
		line := ReadLine(lineNumber)
		fmt.Println(i)
		for !contains(listOfLines, lineNumber) {
			listOfLines = append(listOfLines, lineNumber)
			parts := strings.Split(line, " ")
	
			operation := parts[0]
			argument := parts[1]
			number, err := strconv.Atoi(argument)
			if err != nil {
				// insert error handling here
			  }
			  
			if lineNumber == listOfJmp[i]{
				operation = "nop"
			}

			if operation == "acc" {
				accumulator = accumulator + number 
				if accumulator < 0 {
					accumulator = 0
				}
				lineNumber = lineNumber + 1
			} else if operation == "jmp" {
				lineNumber = lineNumber + number
			} else if operation == "nop" {
				lineNumber = lineNumber + 1
			}
	
			if lineNumber > 630 {
				fmt.Println(lineNumber)
			}
			// fmt.Println(line)
			// fmt.Println(lineNumber)
			// fmt.Println(accumulator)
			if lineNumber == docLength{
				fmt.Println(line)
				fmt.Println(listOfJmp[i])
				fmt.Println("Jmp is changed")
				fmt.Println(accumulator)
			} else {
			line = ReadLine(lineNumber)
			}
		}
	}


	for i := range listOfNop {
		accumulator := 0
		lineNumber := 0
		listOfLines := []int{}
		line := ReadLine(lineNumber)
		fmt.Println(i)
		for !contains(listOfLines, lineNumber) {
			listOfLines = append(listOfLines, lineNumber)
			parts := strings.Split(line, " ")
	
			operation := parts[0]
			argument := parts[1]
			number, err := strconv.Atoi(argument)
			if err != nil {
				// insert error handling here
			  }
			  
			if lineNumber == listOfNop[i]{
				operation = "jmp"
			}

			if operation == "acc" {
				accumulator = accumulator + number 
				if accumulator < 0 {
					accumulator = 0
				}
				lineNumber = lineNumber + 1
			} else if operation == "jmp" {
				lineNumber = lineNumber + number
			} else if operation == "nop" {
				lineNumber = lineNumber + 1
			}
	
			// fmt.Println(len(listOfLines))
			// fmt.Println(docLength)
			// fmt.Println(line)
			// fmt.Println(lineNumber)
			// fmt.Println(accumulator)
			
			if lineNumber > 630 {
				fmt.Println(lineNumber)
			}

			if lineNumber == docLength{
				fmt.Println(line)
				fmt.Println(listOfNop[i])
				fmt.Println("Nop is changed")
				fmt.Println(accumulator)
			}
			line = ReadLine(lineNumber)

		}
	}
	

	
	//Solution 1134
	//Solution part 2 1205
}
