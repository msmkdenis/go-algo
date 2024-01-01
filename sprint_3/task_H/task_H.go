package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Вечером ребята решили поиграть в игру «Большое число».
Даны числа. Нужно определить, какое самое большое число можно из них составить.
*/

func main() {

	const maxCapacity = 5 * 1024 * 1024
	buf := make([]byte, maxCapacity)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)

	var s []string
	for scanner.Scan() {
		line := scanner.Text()
		s = append(s, line)
		if line == "" {
			break
		}
	}
	capacity, _ := strconv.Atoi(s[0])
	unsorted := makeIntSlice(s[1], capacity)
	res := bubbleSort(unsorted, isBigger)
	var answer strings.Builder
	for _, v := range res {
		answer.WriteString(strconv.Itoa(v))
	}
	fmt.Println(answer.String() + "\n")
}

func makeIntSlice(s string, cap int) []int {
	result := make([]int, 0, cap)
	for _, v := range strings.Split(s, " ") {
		i, _ := strconv.Atoi(v)
		result = append(result, i)
	}
	return result
}

func bubbleSort(arr []int, bigger func(int, int) bool) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if !bigger(arr[j], arr[j+1]) {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func isBigger(firstNumber int, secondNumber int) bool {
	s1 := strconv.Itoa(firstNumber)
	s2 := strconv.Itoa(secondNumber)
	sum1 := s1 + s2
	sum2 := s2 + s1
	stringNumber1, _ := strconv.Atoi(sum1)
	stringNumber2, _ := strconv.Atoi(sum2)
	return stringNumber1 > stringNumber2
}
