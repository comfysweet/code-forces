package utils

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetSlices() (int, []int) {
	a := make([]string, 0)
	v := make([]int, 0)

	a = getSlice(a)
	n, err := strconv.Atoi(a[0])
	if err != nil {
		panic(err)
	}
	if n < 0 || n > 1000 {
		panic("n must be in 1≤n≤1000")
	}
	val := strings.Split(a[1], " ")
	for _, s := range val {
		t, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		v = append(v, t)
	}
	if n != len(val) {
		panic("n != len()")
	}
	return n, v
}

func getSlice(a []string) []string {
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < 2; i++ {
		scanner.Scan()
		text := scanner.Text()
		if len(text) != 0 {
			a = append(a, text)
		}
	}
	if scanner.Err() != nil {
		fmt.Println("Error: ", scanner.Err())
	}
	if len(a) == 0 {
		panic(errors.New("text is empty"))
	}
	return a
}
