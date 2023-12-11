package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// https://contest.yandex.ru/contest/22779/problems/L/

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var s []string
	for scanner.Scan() {
		line := scanner.Text()
		s = append(s, line)
		if line == "" {
			break
		}
	}

	inputDigits := strings.Split(s[0], " ")
	n, _ := strconv.Atoi(inputDigits[0])
	k, _ := strconv.Atoi(inputDigits[1])

	ab := []int{1, 1}
	d := math.Pow(10, float64(k))
	fib := 0

	if n <= 2 {
		fib = 1
	} else {
		n -= 1
		for i := 0; i < n; i++ {
			s := (ab[0] + ab[1]) % int(d)
			ab[0] = ab[1]
			ab[1] = s
			fib = ab[1]
		}
	}

	fmt.Println(fib)
}
