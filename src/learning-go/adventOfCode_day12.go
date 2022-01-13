package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func main() {

	f, _ := os.Open("day12.txt")

	horizontal := 0
	vertical := 0

	// lineCount := 0

	// currentDirection := "E"

	waypointH := 10
	waypointV := 1

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		//length := len([]rune(line))
		direction := string([]rune(line)[0])
		digits := string([]rune(line)[1:])
		distance, err := strconv.Atoi(digits)
		if err != nil {
			// insert error handling here
		  }

		if direction == "N" || direction == "S" || direction == "E" || direction == "W" {
			// currentDirection = string([]rune(direction)[0])
			if direction == "N" {
				waypointV = waypointV + distance
			}
	
			if direction == "S" {
				waypointV = waypointV - distance
			}
	
			if direction == "W" {
				waypointH = waypointH - distance
			}
	
			if direction == "E" {
				waypointH = waypointH + distance
			}
		} else {
			if direction == "F" {
				horizontal = horizontal + (waypointH * distance)
				vertical = vertical + (waypointV * distance)
			} else {
				if (direction == "R" && distance == 90) || (direction == "L" && distance == 270) {
					storeValue := waypointH
					fmt.Println(storeValue)
					waypointH = waypointV
					waypointV = -storeValue
					fmt.Println(storeValue)
					fmt.Println(waypointH)
					fmt.Println(waypointV)
				} else if distance == 180 {
					waypointH = -waypointH
					waypointV = -waypointV
				} else {
					storeValue := waypointH
					waypointH = -waypointV
					waypointV = storeValue
				}
			}
			// if currentDirection == "N" {
			// 	if direction == "F" {
			// 		vertical = vertical + (distance * 10)
			// 	} else {
			// 		if (direction == "L" && distance == 90) || (direction == "R" && distance == 270){
			// 			currentDirection = "W"
			// 		} else if distance == 180 {
			// 			currentDirection = "S"
			// 		} else {
			// 			currentDirection = "E"
			// 		}
			// 	} 
			// } else if currentDirection == "S" {
			// 	if direction == "F" {
			// 		vertical = vertical - distance
			// 	} else {
			// 		if (direction == "L" && distance == 90) || (direction == "R" && distance == 270){
			// 			currentDirection = "E"
			// 		} else if distance == 180 {
			// 			currentDirection = "N"
			// 		} else {
			// 			currentDirection = "W"
			// 		}
			// 	}
			// } else if currentDirection == "E" {
			// 	if direction == "F" {
			// 		horizontal = horizontal + distance
			// 	} else {
			// 		if (direction == "L" && distance == 90) || (direction == "R" && distance == 270){
			// 			currentDirection = "N"
			// 		} else if distance == 180 {
			// 			currentDirection = "W"
			// 		} else {
			// 			currentDirection = "S"
			// 		}
			// 	}
			// } else if currentDirection == "W" {
			// 	if direction == "F" {
			// 		horizontal = horizontal - distance
			// 	} else {
			// 		if (direction == "L" && distance == 90) || (direction == "R" && distance == 270){
			// 			currentDirection = "S"
			// 		} else if distance == 180 {
			// 			currentDirection = "E"
			// 		} else {
			// 			currentDirection = "N"
			// 		}
			// 	}
			// }
		}

		fmt.Println(line)
		// fmt.Println(currentDirection)
		fmt.Println(horizontal)
		fmt.Println(vertical)

		// fmt.Println(currentDirection)
		// fmt.Println(direction)
		// fmt.Println(distance)

	}

	fmt.Println(horizontal)
	fmt.Println(vertical)
	if horizontal < 0 && vertical < 0 {
		fmt.Println(horizontal + vertical)
	} else if horizontal < 0 {
		fmt.Println(vertical - horizontal)
	} else if vertical < 0 {
		fmt.Println(horizontal - vertical)
	} else {
		fmt.Println(horizontal + vertical)
	}

	//Solution 1 1007
	//Solution 2 41212
}
