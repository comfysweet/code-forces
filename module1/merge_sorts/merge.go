package main

import (
	"fmt"
	"github.com/comfysweet/code-forces/module1/utils"
)

/**
Сортировка слиянием
Входные данные
Первая строка содержит число 𝑛 (1≤𝑛≤105). Далее идет 𝑛 целых чисел, не превосходящих по абсолютной величине 109.
5
5 4 3 2 1
Выходные данные
Выведите числа в порядке неубывания.
1 2 3 4 5
*/

func main() {
	_, v := utils.GetSlices()

	result := mergeSort(v)
	fmt.Print(result)
}

func mergeSort(v []int) []int {
	n := len(v)
	switch {
	case n > 2:
		lb := mergeSort(v[:n/2])
		rb := mergeSort(v[n/2:])
		v = append(lb, rb...)

		for i := 0; i < n-1; i++ {
			mv := v[i]
			mi := i

			for j := i + 1; j < n; j++ {
				if mv > v[j] {
					mv = v[j]
					mi = j
				}
			}

			if mi != i {
				v[i], v[mi] = v[mi], v[i]
			}
		}

	case len(v) > 1 && v[0] > v[1]:
		v[0], v[1] = v[1], v[0]
	}
	return v
}
