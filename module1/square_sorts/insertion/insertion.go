package main

import (
	"fmt"
	"github.com/comfysweet/code-forces/module1/utils"
)

/**
Сортировка вставками
Входные данные
В первой строке дано целое число 𝑛 (1≤𝑛≤1000) — размер массива.
Во второй строке даны 𝑛 целых чисел, каждое из которых по модулю не превосходит 1000.
5
5 4 3 2 1
Выходные данные
Выведите отсортированный по неубыванию массив.
1 2 3 4 5
*/

func main() {
	n, v := utils.GetSlices()

	result := insertionSort(n, v)
	fmt.Print(result)
}

func insertionSort(n int, v []int) []int {
	for j := 1; j < n; j++ {
		key := v[j]
		i := j - 1
		for i >= 0 && v[i] > key {
			v[i+1] = v[i]
			i = i - 1
		}
		v[i+1] = key
	}
	return v
}
