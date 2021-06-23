package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/**
Предположим, что определено следующее:

[a,b] – отрезок поиска решения

ans – существующее решение на [a,b]

less(x) – функция, соответствующая тому, что x меньше решения (x < ans)
*/

func main() {
	values := getSlices()
	left := values[0]
	right := values[1]
	ans := values[2]

	for left < right {
		x := (left + right) / 2
		if x < ans {
			left = x + 1
		} else {
			right = x
		}
	}
	ans = left
	fmt.Println(ans)
}

func getSlices() []int {
	values := make([]int, 0)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	if len(text) == 0 {
		panic("string is empty")
	}
	s := strings.Split(text, " ")
	for _, s2 := range s {
		v, err := strconv.Atoi(s2)
		if err != nil {
			panic("error values")
		}
		values = append(values, v)
	}
	return values
}
