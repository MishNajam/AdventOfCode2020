package main

import (
	"fmt"
	"bufio"
	"os"
	// "strconv"
)

func getSize() (int,int){
	f, _ := os.Open("day11.txt")
	scanner := bufio.NewScanner(f)
	x := 0 //93 //10
	y := 0 //93 //10
	for scanner.Scan() {
		line := scanner.Text()
		for element := range line {
			if x < element + 1{
				x = element + 1
			}
		}
		y = y + 1
	}
	return x, y
}

func getCount(array []string, character string) int {
	count := 0
	for _,element := range array {
		if element == character{
			count = count + 1
		}
	}
	return count
}

func getOccupiedSeats(seatingState [][]string) int {
	sum := 0
	for y := range seatingState {
		sum = sum + getCount(seatingState[y], "#")
	}
	return sum
}

func createSeatingPlan(number int) [][]string{
	f, _ := os.Open("day11.txt")
	scanner := bufio.NewScanner(f)
	var plan [][] string
	x := 0
	plan = make([][]string, number)
	for scanner.Scan() {
		line := scanner.Text()
		plan[x] = make([]string, len(line))
		for y := range line {
			plan[x][y] = string([]rune(line)[y])
		}
		x = x + 1
	}
	return plan
}

func copySeatingState(number int, seatingState [][]string) [][]string{
	plan := make([][]string, number)
	for y, row := range seatingState {
		plan[y] = make([]string, len(row))
		for x := range row {
			plan[y][x] = seatingState[y][x]
		}
	}
	return plan
}

func main() {
	column, row := getSize()
	seatingPlan := createSeatingPlan(column)
	seatingState := createSeatingPlan(column)
	changeMade := true

	fmt.Println(column)
	fmt.Println(row)

	loopCount := 0

	for changeMade {
		//canner := bufio.NewScanner(f)
		changeMade = false
		
		x := 0
		y := 0

		for _, yrow := range seatingPlan {
			for _, value := range yrow {
				seat := value
				seats := make([]string,0)

				offset := 1
				for y-offset >= 0 {
					north := seatingPlan[y-offset][x]
					if north == "#" || north == "L" {
						seats = append(seats, north)
						break
					}
					offset = offset + 1
				}

				offset = 1
				for y+offset < row {
					south := seatingPlan[y+offset][x]
					if south == "#" || south == "L" {
						seats = append(seats, south)
						break
					}
					offset = offset + 1
				}

				offset = 1
				for x+offset < column {
					east := seatingPlan[y][x+offset]
					if east == "#" || east == "L" {
						seats = append(seats, east)
						break
					}
					offset = offset + 1
				}

				offset = 1
				for x-offset >= 0 {
					west := seatingPlan[y][x-offset]
					if west == "#" || west == "L" {
						seats = append(seats, west)
						break
					}
					offset = offset + 1
				}

				offset = 1
				for y-offset >= 0 && x-offset >= 0 {
					nw := seatingPlan[y-offset][x-offset]
					if nw == "#" || nw == "L" {
						seats = append(seats, nw)
						break
					}
					offset = offset + 1
				}

				offset = 1
				for y-offset >= 0 && x+offset < column {
					ne := seatingPlan[y-offset][x+offset]
					if ne == "#" || ne == "L" {
					seats = append(seats, ne)
					break
					}
					offset = offset + 1
				}

				offset = 1
				for y+offset < row && x-offset >= 0 {
					sw := seatingPlan[y+offset][x-offset]
					if sw == "#" || sw == "L" {
					seats = append(seats, sw)
					break
					}
					offset = offset + 1
				}

				offset = 1
				for y+offset < row && x+offset < column {
					se := seatingPlan[y+offset][x+offset]
					if se == "#" || se == "L" {
					seats = append(seats, se)
					break
					}
					offset = offset + 1
				}

				//fmt.Println(seats)

				if seat == "L" { // empty
					if getCount(seats, "#") == 0 {
						seatingState[y][x] = "#"
						changeMade = true
					}
				} else if seat == "#" { //occupied
					if getCount(seats, "#") >= 5 {
						seatingState[y][x] = "L"
						changeMade = true
					}
				} else if seat == "." { //floor
					seatingState[y][x] = "."
				}

				x = (x + 1) % column
			}
			y = y + 1
		}

		loopCount = loopCount + 1
	//	for scanner.Scan() {
		//	line := scanner.Text()
			// length := len([]rune(line))

			// for xcoord := range line {
				// seat := string([]rune(line)[xcoord])

				//fmt.Println(seats)

				
			// }	
			// fmt.Println(length)
		// }
	 
	// }
	fmt.Println(seatingState)
	fmt.Println(seatingPlan)
	

	fmt.Println(column)
	seatingPlan = copySeatingState(column, seatingState)

	}

	occupiedSeats := getOccupiedSeats(seatingPlan)
	fmt.Println(occupiedSeats)//Solution 1 2249 //Solution 2 2023
	
}
