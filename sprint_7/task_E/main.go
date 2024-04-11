package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputData := getInputData()
	x, _ := strconv.Atoi(inputData[0])
	k, _ := strconv.Atoi(inputData[1])
	banknotes := parseBanknotes(inputData[2], k)

	result := minBanknotes(x, banknotes)
	fmt.Println(result)
}

func minBanknotes(x int, banknotes []int) int {
	const INF = 1e9 + 7
	dp := make([]int, x+1)

	for i := 1; i <= x; i++ {
		dp[i] = INF
		for _, note := range banknotes {
			if i-note >= 0 && dp[i-note]+1 < dp[i] {
				dp[i] = dp[i-note] + 1
			}
		}
	}

	if dp[x] == INF {
		return -1 // Если невозможно набрать сумму x
	}

	return dp[x]
}

func parseBanknotes(input string, k int) []int {
	banknotes := make([]int, k)
	parts := strings.Fields(input)
	for i := 0; i < k; i++ {
		banknotes[i], _ = strconv.Atoi(parts[i])
	}
	return banknotes
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
