package main

import "fmt"

//exercise-fibonacci-closure.go
// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	var pre, cur = 0, 0
	return func() int {
		temp := cur
		cur = pre + cur
		pre = temp
		if cur == 0 {
			cur = 1
			return 0
		}
		return temp
	}
}
func fibonacci1() func() int {
	var pre, cur, idx = 0, 1, -1
	return func() int {
		idx++
		if idx <= 1 {
			return idx
		}
		temp := cur
		cur = pre + cur
		pre = temp
		return cur
	}
}
func main() {
	f := fibonacci()
	f1 := fibonacci1()
	for i := 0; i < 10; i++ {
		fmt.Println(f(), " = ", f1())
	}
}
