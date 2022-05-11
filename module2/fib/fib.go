package main

import "fmt"

var F = make(map[int]int)

func main() {
	var n int
	fmt.Scanln(&n)
	f := getFibNub(n)
	fmt.Println(f)
}

func getFibNub(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	if v, ok := F[n]; ok == true {
		return v
	}
	F[n] = getFibNub(n-1) + getFibNub(n-2)
	return F[n]
}
