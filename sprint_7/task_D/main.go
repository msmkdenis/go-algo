package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	inputData := getInputData()
	n, _ := strconv.Atoi(inputData[0])

	fmt.Println(fibonacci(n))
}

func fibonacci(n int) int {
	const mod = 1e9 + 7
	if n == 0 || n == 1 {
		return 1
	}

	prev, curr := 1, 1
	for i := 2; i <= n; i++ {
		next := (prev + curr) % mod
		prev = curr
		curr = next
	}

	return curr
}

func getInputData() []string {
	input, _ := os.Open("input.txt")
	defer input.Close()

	const maxCapacity = 10240 * 1024
	buf := make([]byte, maxCapacity)

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)
	scanner.Buffer(buf, maxCapacity)

	var s []string
	for scanner.Scan() {
		bufStr := scanner.Text()
		s = append(s, bufStr)
	}

	return s
}
