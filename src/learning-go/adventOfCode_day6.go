package main

import (
	"fmt"
	"bufio"
	"os"
)

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func sum(numbers []int) int {
	var sum = 0
	for i := 0; i < len(numbers); i++ {
		sum = sum + numbers[i]		
	}
	return sum
}

func remove(s []string, i int) []string {
	if i == len(s){
		s = s[:i]
	} else {
		s = append(s[:i], s[i+1:]...)
	}
    return s
}

func indexOf(element string, data []string) (int) {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1    //not found.
 }

func main() {

	f, _ := os.Open("day6.txt")

	characters := []string{}
	questions := []int{}
	similar := []string{}

	lineCount := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		length := len([]rune(line))

		if length == 0 {
			questions = append(questions, len(similar))
			characters = []string{}
			similar = []string{}
		} else if len(characters) < 1 {
			for i:= range line {
				character := string([]rune(line)[i])
				characters = append(characters, character)	
			}
			similar = characters
		} else {
			characters = []string{}
			for i:= range line {
				character := string([]rune(line)[i])
				characters = append(characters, character)
			}

			cloneSimilar := []string{}
			for i:= range similar {
				character := similar[i]
				cloneSimilar = append(cloneSimilar, character)
			}
			
			if lineCount == 7 || lineCount == 8 {
				fmt.Println(similar)
				fmt.Println(cloneSimilar)
			}
				for i:= range similar {
					character := similar[i]
					if lineCount == 7 || lineCount == 8 {
						fmt.Println(character)
					}
					if !contains(characters, character) {
						index := indexOf(character, cloneSimilar)  
						if contains(cloneSimilar, character)  {
							cloneSimilar = remove(cloneSimilar, index)
						}
					}
					// fmt.Printf("IN LOOP Clone of similar is: %v\n", cloneSimilar)
				}

				similar = cloneSimilar
			// fmt.Printf("End of loop similar is: %v\n", similar)
			// fmt.Println(characters)
		}

		// fmt.Printf("line %v\n", line)
		// fmt.Printf("length %v\n", length)
		// fmt.Printf("questions %v\n", questions)
		// fmt.Printf("similar letters %v\n", similar)
		// fmt.Printf("Line %v has %v similar letters %v\n", lineCount, len(similar), similar)
		lineCount++
		}
		questions = append(questions, len(similar))
		fmt.Printf("questions %v\n", questions)


		fmt.Println(sum(questions))

		//Solution 1: 6596
		//Solution 2 : 3535 (wrong)
		//Solution 2 : 3353 (wrong)
		//Solution 2 : 3219 (correct!)
	}
