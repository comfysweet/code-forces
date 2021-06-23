package main

import "fmt"

//Вычисление квадратного корня
func main() {
	var l, r, a, mid float64
	eps := 0.001
	fmt.Scanln(&a)
	l = 0
	r = a + 1
	for r-l > eps {
		mid = (l + r) / 2
		if (mid * mid) < a {
			l = mid
		} else {
			r = mid
		}
	}
	res := fmt.Sprintf("%.2f", r)
	fmt.Println(res)
}
