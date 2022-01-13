package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func summed(numbers []int) int {
	var sum = 0
	for i := 0; i < len(numbers); i++ {
		sum = sum + numbers[i]		
	}
	return sum
}

func min(numbers []int) int {
	var min = 99999999999
	for i := 0; i < len(numbers); i++ {
		if numbers[i] < min {
			min = numbers[i]
		}
	}
	return min
}

func max(numbers []int) int {
	var max = 0
	for i := 0; i < len(numbers); i++ {
		if numbers[i] > max {
			max = numbers[i]
		}	
	}
	return max
}

func ReadLine(lineNum int) (line string) {
	f, _ := os.Open("day9.txt")
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

func findSum(number int) (numbers []int) {
	f, _ := os.Open("day9.txt")
	sc := bufio.NewScanner(f)
	
	loop := 0
	preamble := []int{}
    for sc.Scan() {

		sum := 0
		preamble = []int{}

		lineCount := loop
		//linenumber is 585
		for sum < number {
			loopLine := ReadLine(lineCount)
			loopdigit, err := strconv.Atoi(loopLine)
			if err != nil {
				// insert error handling here
			}
			preamble = append(preamble, loopdigit)
			sum = summed(preamble)
			if sum == number && len(preamble) > 1{
				return preamble
				break
			}
			fmt.Println(sum)
			fmt.Println(lineCount)
			lineCount = lineCount + 1
		}
		loop = loop + 1
	}
	
    return preamble
}

func main() {

	f, _ := os.Open("day9.txt")

	scanner := bufio.NewScanner(f)
	lineCount := 0
	preamble := []int{}
	counter := 0 
	numberToSum := 0

	for scanner.Scan() {
		line := scanner.Text()
		number, err := strconv.Atoi(line)
		if err != nil {
			// insert error handling here
		}
		if lineCount < 25 {
			preamble = append(preamble, number)
		} else {
			flag := false
			for i := 0; i < len(preamble); i++ {
				for j := i+1 ; j < len(preamble); j++ {
					if (preamble[i] + preamble[j] == number){
						flag = true
					}
				}
			}

			if !flag {
				fmt.Println(lineCount)
				fmt.Println(number)
				numberToSum = number
				break
			} else {
				preamble[counter] = number
				// fmt.Println(preamble[counter])
				// fmt.Println(preamble)
				counter = (counter+1)%25
			}
		}




		// fmt.Println(line)
		// fmt.Println(preamble)

		lineCount = lineCount + 1
	}


	one := findSum(numberToSum)
	fmt.Println(one)
	min := min(one)
	max := max(one)
	fmt.Println(min)
	fmt.Println(max)

	fmt.Println(min + max)

	//solution 1: 167829540
}
