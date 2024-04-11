package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := getInputData()
	s := input[0]
	t := input[1]

	fmt.Println(LevenshteinDistance(s, t))
}

func LevenshteinDistance(s, t string) int {
	n, m := len(s), len(t)

	// Создаем двумерный массив размером (n+1) x (m+1) для хранения расстояний Левенштейна.
	// dp[i][j] будет содержать расстояние Левенштейна для подстрок s[:i] и t[:j].
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}

	// Инициализируем значения в первой строке и первом столбце.
	for i := 0; i <= n; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= m; j++ {
		dp[0][j] = j
	}

	// Заполняем массив dp.
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			// Если текущие символы в строках s и t совпадают, то расстояние Левенштейна не изменяется.
			// В противном случае, берем минимум из трех возможных операций: вставки, удаления и замены.
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1])
			}
		}
	}

	// Расстояние Левенштейна для всей строки s и t находится в правом нижнем углу массива dp.
	return dp[n][m]
}

func min(a, b, c int) int {
	minVal := a
	if b < minVal {
		minVal = b
	}
	if c < minVal {
		minVal = c
	}
	return minVal
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
