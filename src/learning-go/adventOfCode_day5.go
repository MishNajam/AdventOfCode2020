package main

import (
	"fmt"
	"bufio"
	"os"
)
//FFBBBFB
func calculateRow(text string) int {
	topRange:= 0
	bottomRange:= 127
	for i := range text {
		if string([]rune(text)[i]) == "F" {
			topRange= topRange //0 ///28
			bottomRange= topRange + ((bottomRange-topRange)/2) //1 ->64, 2->32, //30
		} else if string([]rune(text)[i]) == "B" {
			topRange= bottomRange - ((bottomRange-topRange)/2) //24 5th 28 29
			bottomRange= bottomRange //32->32 //30
		}
	}
	return bottomRange 
}

func calculateColumn(text string) int {
	leftRange:= 0
	rightRange:= 7
	for i:= range text {//evaluates to 1
		if string([]rune(text)[i]) == "L"{
			leftRange= leftRange 
			rightRange= leftRange + ((rightRange-leftRange)/2)
		} else if string([]rune(text)[i]) == "R" {
			leftRange= rightRange - ((rightRange-leftRange)/2)
			rightRange= rightRange 
		}
	}
	return leftRange 
   }

func calculateSeatId(row int, column int) int {
	calculation := (row * 8) + column
	return calculation
}

func main() {

	f, _ := os.Open("day5.txt")

	max := 0

	seats := [8][128]int{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		// length := len([]rune(line))

		rowString := line[0:7]
		columnString := line[7:10] 

		row := calculateRow(rowString)
		column := calculateColumn(columnString)
		seatId := calculateSeatId(row, column)

		seats[column][row] = 1

			if max < seatId {
				max = seatId
			}

			fmt.Println(line)
			// fmt.Println(length)
			fmt.Println(row)
			fmt.Println(column)
			fmt.Println(seatId)
			fmt.Println(max)
		}

	
		fmt.Println(max)
		//Solution 1 : 816
		fmt.Println(seats)
		for  i := 0; i < 8; i++ {
			for j := 4; j <= 120; j++ {
				if seats[i][j] == 0 && seats[i][j+1] == 1 && seats[i][j-1] == 1{
				// if seats[i-1][j] == 1 && seats[i][j] == 0 && seats[i+1][j]{
				fmt.Println(i)
				fmt.Println(j)
				fmt.Println(calculateSeatId(j,i))
				}
			}
		 }
	}
