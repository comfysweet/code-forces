package main

import (
	"fmt"
	"github.com/comfysweet/code-forces/module1/utils"
)

/**
Сортировка выбором
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

	result := selectionSort(n, v)
	fmt.Print(result)
}

func selectionSort(n int, v []int) []int {
	for i := 0; i < n-1; i++ {
		least := i
		for j := i + 1; j < n; j++ {
			if v[j] < v[least] {
				least = j
			}
		}
		v[i], v[least] = v[least], v[i]
	}
	return v
}
