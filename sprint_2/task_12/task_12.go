package main

import (
	"fmt"
	"strconv"
)

// https://contest.yandex.ru/contest/22779/problems/L/

func main() {

	var input string
	fmt.Scan(&input)

	number, _ := strconv.Atoi(input)
	fmt.Println(Solution(number + 1))
}

func Solution(number int) int {
	if number == 1 || number == 2 {
		return 1
	}

	return Solution(number-1) + Solution(number-2)
}
