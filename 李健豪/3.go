package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	var year, month, day string

	fmt.Println("What year were you born [YYYY]?")
	fmt.Scanln(&year)

	fmt.Println("What month were you born [MM]?")
	fmt.Scanln(&month)

	fmt.Println("What day were you born [DD]?")
	fmt.Scanln(&day)

	fmt.Printf("It looks like you were born on " + day + " / " + month + " / " + year)

	var minus int

	now := time.Now()

	local, _ := time.LoadLocation("Local")

	t, _ := time.ParseInLocation("2006-01-02 15:04:05", strconv.Itoa(time.Now().Year())+"-"+month+"-"+day+" 00:00:00", local)

	minus = SubDays(now, t)

	if minus == 0 {
		fmt.Println("\n", "Aha,today is your birthday!", "\n", " Happy birthday to you!")
	} else if minus > 0 {
		fmt.Println("\n", "It looks like your birthday has passed "+strconv.Itoa(minus)+" days.", "\n", "Hope you're looking forward to the next one!")
	} else {
		fmt.Println("\n", "It looks like there are "+strconv.Itoa(-minus)+" days from your birthday.", "\n", " Hope you're looking forward to it!")
	}
}
func SubDays(t1, t2 time.Time) (day int) {
	minus := (t1.Sub(t2)).Hours()
	if minus < 0 {
		day = int(minus-24) / 24
	} else {
		day = int(minus / 24)
	}
	return day
}
