// print out 10 random dates
package main

import (
	"fmt"
	"math/rand"
)

func IsLeapYear(year int) bool {
	if year % 400 == 0 || (year % 100 != 0 && year % 4 == 0) {
		return true 
	} else {
		return false
	}
}

var era = "AD"

func Date(randomYear int) (int, int, int)  {
	year := randomYear
	month := rand.Intn(12) + 1
	daysInMonth := 31 

	switch month {
	case 2:
		if IsLeapYear(year) {
			daysInMonth = 29
		} else {
			daysInMonth = 28
		}
	case 4, 6, 9, 11:
		daysInMonth = 30
	}
	day := rand.Intn(daysInMonth) + 1
	return  year, month ,day 
}


func main() {	
	for count := 0; count < 10; count++ {
		var randomYear = rand.Intn(21) +1 + 1999
		year, month, day := Date(randomYear)
		fmt.Println(era, year, month, day)		
	}	
}
