package main

import "fmt"

// 返回一个“返回int的函数”
func fibonacci() func() int {
	x1 := 0
	x2 := 1
	return func() int {
		x1, x2 = x2, x1+x2
		return x2 - x1
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
