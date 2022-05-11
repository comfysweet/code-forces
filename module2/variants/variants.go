package main

import "fmt"

/**
Требуется подсчитать количество последовательностей длины N , состоящих из 0 и 1, в которых никакие две единицы не стоят рядом.
*/

func main() {
	var n int
	fmt.Scanln(&n)
	f := make([][]int, 2, 2)
	for i := range f {
		f[i] = make([]int, n, n)
	}
	//ff := make([]int, 0)

	if n == 1 {
		f[0][n-1] = 1
		f[1][n-1] = 1
	} else {
		for i := 2; i < n+1; i++ {
			f[0][i] = f[0][i-1] + f[1][i-1]
			f[1][i] = f[0][i-1]
		}
	}
	res := f[0][n-1] + f[1][n-1]
	fmt.Println(res)
}

func fuc(f [][]int, n int) {
	for i := 2; i < n+1; i++ {
		f[0][i] = f[0][i-1] + f[1][i-1]
		f[1][i] = f[0][i-1]
	}
}