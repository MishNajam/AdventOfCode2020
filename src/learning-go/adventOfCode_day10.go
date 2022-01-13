package main

import (
	"fmt"
	"bufio"
	"os"
	"sort"
	"strconv"
)

// func contains(s []string, str string) bool {
// 	for _, v := range s {
// 		if v == str {
// 			return true
// 		}
// 	}

// 	return false
// }

func findMax() (max int) {
	f, _ := os.Open("day10.txt")
	sc := bufio.NewScanner(f)
	max = 0
    for sc.Scan() {
		line := sc.Text()
		number, err := strconv.Atoi(line)
		if err != nil {
			// insert error handling here
		} 
        if number > max {
            max = number
		}
    }
    return max
}

func getNumbers() ([]int) {
	f, _ := os.Open("day10.txt")
	sc := bufio.NewScanner(f)
	numbers := []int{}
    for sc.Scan() {
		line := sc.Text()
		number, err := strconv.Atoi(line)
		if err != nil {
			// insert error handling here
		} 
        numbers = append(numbers, number)
    }
    return numbers
}

// for i := 0; i < len(numbers); i++ {


// func variations(numbers []int) int {
// 	pathCount := [len(numbers)]int{}
// 	pathCount[0] = 1
// 	for i := 1; i < len(numbers); i++ {
// 		for j := i - 3; j < i; j++ {
// 			if j >= 0 && numbers[i] - numbers[j] <= 3 {
// 				pathCount[i] = pathCount[i] + pathCount[j]
// 			}
// 		}
// 	}
// 	return pathCount[len(pathCount) - 1]
// }

func main() {

	// f, _ := os.Open("day10sample.txt")

	buffer := 3
	max := findMax()
	myVoltage := max + buffer
	chargingOutletVoltage:= 0

	oneJoltCount := 0
	twoJoltCount := 0
	threeJoltCount := 0
	oneVolt := []int{}
	twoVolt := []int{}
	threeVolt := []int{}

	numbers := getNumbers()
	fmt.Println(numbers)
	// numbers = append(numbers,chargingOutletVoltage)
	// numbers = append(numbers,myVoltage)

	sort.Ints(numbers)
	fmt.Println(numbers)

	length:= len(numbers) + 2
	var entireList [length]int
	entireList = append(entireList, chargingOutletVoltage)
	entireList = append(entireList, numbers)
	entireList = append(entireList, myVoltage)

	for i := 0; i < len(numbers); i++ {
		subtractor := 0

		if i == 0 {
			subtractor = chargingOutletVoltage
		} else {
			subtractor = numbers[i-1]
		}
		if numbers[i] - subtractor == 3{
			threeVolt = append(threeVolt, numbers[i])
			threeJoltCount = threeJoltCount + 1
		} else if numbers[i] - subtractor == 2 {
			twoVolt = append(twoVolt, numbers[i])
			twoJoltCount = twoJoltCount + 1
		} else if numbers[i] - subtractor == 1 {
			oneVolt = append(oneVolt, numbers[i])
			oneJoltCount = oneJoltCount + 1
		}
	}

	threeJoltCount = threeJoltCount + 1

	fmt.Println(numbers)
	fmt.Println(myVoltage)
	fmt.Println(oneVolt)
	fmt.Println(threeVolt)

	fmt.Println(oneJoltCount)
	fmt.Println(threeJoltCount)
	fmt.Println(oneJoltCount * threeJoltCount)

	fmt.Println(entireList)
	// fmt.Println(variations(entireList))


	//Solution 1 = 1625
	}
