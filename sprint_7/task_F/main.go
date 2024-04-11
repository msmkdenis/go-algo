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
	n, k := parseInput(inputData[0])

	// Создаем массив dp для хранения количества способов достичь каждой ступени.
	// dp[i] будет содержать количество способов достичь ступени i.
	dp := make([]int, n+1)
	dp[0] = 1 // Начальное количество способов для первой ступени

	// Заполняем массив dp с использованием динамического программирования.
	for i := 1; i <= n; i++ {
		for j := 1; j <= k && j <= i; j++ { // Учитываем ограничение на прыжок до ступени n
			dp[i] = (dp[i] + dp[i-j]) % (1e9 + 7)
		}
	}

	fmt.Println(dp[n-1])
}

func parseInput(input string) (int, int) {
	parts := strings.Fields(input)
	n, _ := strconv.Atoi(parts[0])
	k, _ := strconv.Atoi(parts[1])
	return n, k
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
