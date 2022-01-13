package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"unicode"
)

func validateYear(text string, code string, lowYear int, highYear int) bool {
	if strings.Contains(text, code){
	   parts := strings.Split(text, " ")
	   for i, part := range parts {
		   if strings.Contains(parts[i], code){
			   validate := strings.Split(part, ":")[1]
			   year, err := strconv.Atoi(validate)
			   if err != nil {
				   // insert error handling here
				 }
			   if year >= lowYear && year <= highYear {
				   return true
			   }
		   }
	   }
	} 
	return false
}

func validateHeight(text string) bool {
	if strings.Contains(text, "hgt"){
	   parts := strings.Split(text, " ")
	   for i, part := range parts {
		   if strings.Contains(parts[i], "hgt"){
			   validate := strings.Split(part, ":")[1]
			   if strings.Contains(validate, "cm"){
				height := strings.Split(validate, "cm")[0]
				cm, err := strconv.Atoi(height)
				if err != nil {
					// insert error handling here
				  }
				if cm >= 150 && cm <= 193 {
					return true
				}
			   } else if strings.Contains(validate, "in"){
				height := strings.Split(validate, "in")[0]
				inches, err := strconv.Atoi(height)
				if err != nil {
					// insert error handling here
				  }
				if inches >= 59 && inches <= 76 {
					return true
			   }
			}
		   }
	   }
	} 
	return false
}

func validateHair(text string) bool {
	if strings.Contains(text, "hcl"){
	   parts := strings.Split(text, " ")
	   for i, part := range parts {
		   if strings.Contains(parts[i], "hcl"){
			   validate := strings.Split(part, ":")[1]
				 if len([]rune(validate)) == 7 {
					for i := range validate {
						if i == 0 {
							if string([]rune(validate)[i]) != "#"{
								return false
							}
						} else if i >= 1 {
							if !unicode.IsDigit([]rune(validate)[i]) && !unicode.IsLetter([]rune(validate)[i]) {
								return false
							}
						}
					}
					return true
				 }
		   }
	   }
	} 
	return false
}

func validateEyeColor(text string) bool {
	if strings.Contains(text, "ecl"){
	   parts := strings.Split(text, " ")
	   for i, part := range parts {
		   if strings.Contains(parts[i], "ecl"){
			   validate := strings.Split(part, ":")[1]
			   condition := validate == "amb" || validate == "blu" || validate == "brn" || validate == "gry" || validate == "grn" || validate == "hzl" || validate == "oth"
			   if condition {
					return true
				}
		   }
	   }
	} 
	return false
   }

   func validatePassport(text string) bool {
	if strings.Contains(text, "pid"){
	   parts := strings.Split(text, " ")
	   for i, part := range parts {
		   if strings.Contains(parts[i], "pid"){
			   validate := strings.Split(part, ":")[1]
				 if len([]rune(validate)) == 9 {
					for i := range validate {
						if !unicode.IsDigit([]rune(validate)[i]){
							return false
						}
					}
					return true
				 }
		   }
	   }
	} 
	return false
}

func main() {

	f, _ := os.Open("day4.txt")

	birthYear := false
	issueYear := false
	expirationYear := false
	height := false
	hair := false
	eye := false
	passport := false
	country := false
	//make country passable

	validPassports := 0
	passable := 0
	numberOfPassports := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		length := len([]rune(line))

		if length == 0 {
			numberOfPassports = numberOfPassports + 1
			if birthYear && issueYear &&
			expirationYear && height &&
			hair && eye && passport {
				passable = passable + 1
			}
			if birthYear && issueYear &&
			expirationYear && height &&
			hair && eye && passport && country {
				validPassports = validPassports + 1
			}
			birthYear = false
			issueYear = false
			expirationYear = false
			height = false
			hair = false
			eye = false
			passport = false
			country = false
		} else {
			if birthYear == false {
				birthYear = validateYear(line, "byr", 1920, 2002)
			}
			if issueYear == false {
				issueYear = validateYear(line, "iyr", 2010, 2020)
			}
			if expirationYear == false {
				expirationYear = validateYear(line, "eyr", 2020, 2030)
			}
			if height == false {
				height = validateHeight(line)
			}
			if hair == false {
				hair = validateHair(line)
			}
			if eye == false {
				eye = validateEyeColor(line)
			}
			if passport == false {
				passport = validatePassport(line)
			}
			if country == false {
				country = strings.Contains(line, "cid")
			}
		}

		fmt.Println(line)
		fmt.Println(length)
		fmt.Println(validPassports)
		fmt.Println(passable)
		fmt.Println(numberOfPassports)
	}


	numberOfPassports = numberOfPassports + 1
	if birthYear && issueYear &&
	expirationYear && height &&
	hair && eye && passport {
		passable = passable + 1
	}
	if birthYear && issueYear &&
	expirationYear && height &&
	hair && eye && passport && country {
		validPassports = validPassports + 1
	}

	fmt.Println(validPassports)
	fmt.Println(passable)
	fmt.Println(numberOfPassports)


	//Solution 1: 233
}
