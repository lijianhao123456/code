package main

import (
	"fmt"
	"strconv"
	"time"
)

//by 李健豪
func isrun(i int) bool {
	if (i%100 != 0 && i%4 == 0) || (i%400 == 0) {
		return true
	} else {
		return false
	}
}
func main() {
	var year, month, day int

	a := [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	t1 := time.Now()
	nyear := t1.Year()
	nday := t1.YearDay()
	nday1 := t1.Day()
	sum := 0
	str := t1.Format("2006-01-02")
	strmonth := str[5 : len(str)-3]
	nmonth, _ := strconv.Atoi(strmonth)

Begin:

	fmt.Println("What year were you born?")
	fmt.Scanf("%4d\n", &year)

	if year > nyear {
		println("输入日期不合法")
		goto Begin
	}

	fmt.Println("What month were you born?")
	fmt.Scanf("%2d\n", &month)

	if year == nyear && month > nmonth {
		println("输入日期不合法")
		goto Begin
	}

	fmt.Println("What day were you born?")
	fmt.Scanf("%2d\n", &day)
	if year == nyear && month == nmonth && day > nday1 {
		println("输入日期不合法")
		goto Begin
	}

	fmt.Printf("It looks like you were born on %v / %v / %v\n", day, month, year)

	for i := 0; i < month-1; i++ {
		sum = sum + a[i]
	}
	if isrun(year) && month > 2 {
		sum = sum + 1
	}
	sum += day

	var minus = 0 //下面其实就第二种情况输出比较特殊 其实可以把其他的输出全放一起
	if sum <= 59 && nday <= 59 {
		if sum > nday {
			minus = sum - nday
			fmt.Println("It looks like there are " + strconv.Itoa(minus) + " days from your birthday.")
			fmt.Println("Hope you're looking forward to the next one!")
		} else if sum == nday {
			fmt.Println("Aha,today is your birthday!")
			fmt.Println("Happy birthday to you!")
		} else {
			minus = nday - sum
			fmt.Println("It looks like your birthday has passed " + strconv.Itoa(minus) + " days.")
			fmt.Println("Hope you're looking forward to the next one!")
		}
	} else if isrun(year) && sum == 60 {
		if !isrun(nyear) {
			minus = 366 - nday + sum
			for i := nyear + 1; i <= nyear+3; i++ {
				if isrun(i) {
					break
				} else {
					minus = minus + 365
				}
			}
			fmt.Println("It looks like there are " + strconv.Itoa(minus) + " days from your birthday.")
			fmt.Println("Hope you're looking forward to the next one!")
		} else {
			if sum > nday {
				minus = sum - nday
				fmt.Println("It looks like there are " + strconv.Itoa(minus) + "days from your birthday.")
				fmt.Println("Hope you're looking forward to the next one!")
			} else if sum == nday {
				fmt.Println("Aha,today is your birthday!")
				fmt.Println("Happy birthday to you!")
			} else {
				minus = nday - sum
				fmt.Println("It looks like your birthday has passed " + strconv.Itoa(minus) + " days.")
				fmt.Println("Hope you're looking forward to the next one!")
			}
		}
	} else if isrun(year) && isrun(nyear) {
		if sum > nday {
			minus = sum - nday
			fmt.Println("It looks like there are " + strconv.Itoa(minus) + " days from your birthday.")
			fmt.Println("Hope you're looking forward to the next one!")
		} else if sum == nday {
			fmt.Println("Aha,today is your birthday!")
			fmt.Println("Happy birthday to you!")
		} else {
			minus = nday - sum
			fmt.Println("It looks like your birthday has passed " + strconv.Itoa(minus) + " days.")
			fmt.Println("Hope you're looking forward to the next one!")
		}
	} else if isrun(year) && !isrun(nyear) {
		sum = sum - 1
		if sum > nday {
			minus = sum - nday
			fmt.Println("It looks like there are " + strconv.Itoa(minus) + " days from your birthday.")
			fmt.Println("Hope you're looking forward to the next one!")
		} else if sum == nday {
			fmt.Println("Aha,today is your birthday!")
			fmt.Println("Happy birthday to you!")
		} else {
			minus = nday - sum
			fmt.Println("It looks like your birthday has passed " + strconv.Itoa(minus) + " days.")
			fmt.Println("Hope you're looking forward to the next one!")
		}
	} else if !isrun(year) && isrun(nyear) {
		nday = nday - 1
		if sum > nday {
			minus = sum - nday
			fmt.Println("It looks like there are " + strconv.Itoa(minus) + " days from your birthday.")
			fmt.Println("Hope you're looking forward to the next one!")
		} else if sum == nday {
			fmt.Println("Aha,today is your birthday!")
			fmt.Println("Happy birthday to you!")
		} else {
			minus = nday - sum
			fmt.Println("It looks like your birthday has passed " + strconv.Itoa(minus) + " days.")
			fmt.Println("Hope you're looking forward to the next one!")
		}

	} else if !isrun(year) && isrun(nday) {
		if sum > nday {
			minus = sum - nday
			fmt.Println("It looks like there are " + strconv.Itoa(minus) + " days from your birthday.")
			fmt.Println("Hope you're looking forward to the next one!")
		} else if sum == nday {
			fmt.Println("Aha,today is your birthday!")
			fmt.Println("Happy birthday to you!")
		} else {
			minus = nday - sum
			fmt.Println("It looks like your birthday has passed " + strconv.Itoa(minus) + " days.")
			fmt.Println("Hope you're looking forward to the next one!")
		}
	}
}
