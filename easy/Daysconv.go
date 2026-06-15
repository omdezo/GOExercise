package easy

import "fmt"


func DaysToweeksconv(days int){
	weeks := days/7
	remainDays := days%7
	if remainDays == 0 {
		fmt.Printf("%d days is equal to %d week/s",days,weeks)
	} else {
		fmt.Printf("%d days is equal to %d week/s and %d day/s",days,weeks,remainDays)
	}
	
}