package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/**
Дано 𝑛 (1≤𝑛≤104) отрезков провода длиной 𝑙1, 𝑙2, ..., 𝑙𝑛 (100≤𝑙𝑖≤107) сантиметров.
Требуется с помощью разрезания получить из них 𝑘 (1≤𝑘≤104) равных отрезков как можно большей длины,
выражающейся целым числом сантиметров. Если нельзя получить 𝑘 отрезков длиной даже 1 см, вывести 0.

Входные данные
На первой строке заданы числа 𝑛 и 𝑘. В следующих 𝑛 строках заданы 𝑙𝑖 по одному в строке. Все числа целые.
4 11
802
743
457
539
Выходные данные
Выведите одно число — полученную длину отрезков.
200
 */

func main() {
	_, k, lens := getSlice()

	l := 0
	r := getMax(lens)
	for (r - l) > 0 {
		mid := (l + r) / 2
		avg := 0
		for _, l1 := range lens {
			avg = (l1 / mid) + avg
		}
		if avg >= k {
			l = mid + 1
		} else {
			r = mid
		}
	}
	fmt.Println(l-1)
}

func getSlice() (n int, k int, res []int) {
	val := make([]string, 0)
	res = make([]int, 0)
	var v int
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	text := scan.Text()
	val = strings.Split(text, " ")
	n, _ = strconv.Atoi(val[0])
	k, _ = strconv.Atoi(val[1])
	for i := 0; i < n; i++ {
		fmt.Scanln(&v)
		res = append(res, v)
	}
	return
}

func getMax(in []int) (m int) {
	for i, e := range in {
		if i == 0 || e < m {
			m = e
		}
	}
	return
}
