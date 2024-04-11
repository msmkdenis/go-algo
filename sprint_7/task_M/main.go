package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := getInputData()

	// Чтение количества предметов и грузоподъемности рюкзака
	params := strings.Fields(input[0])
	n, _ := strconv.Atoi(params[0])
	M, _ := strconv.Atoi(params[1])

	// Создание слайсов для хранения весов и значимостей предметов
	weights := make([]int, n)
	values := make([]int, n)

	// Чтение весов и значимостей предметов
	for i := 0; i < n; i++ {
		item := strings.Fields(input[i+1])
		weights[i], _ = strconv.Atoi(item[0])
		values[i], _ = strconv.Atoi(item[1])
	}

	// Создание двумерного слайса для хранения максимальной значимости
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, M+1)
	}

	// Заполнение массива dp
	for i := 1; i <= n; i++ {
		for w := 0; w <= M; w++ {
			dp[i][w] = dp[i-1][w]
			if weights[i-1] <= w {
				dp[i][w] = max(dp[i][w], dp[i-1][w-weights[i-1]]+values[i-1])
			}
		}
	}

	// Восстановление ответа
	res := make([]int, 0)
	i, j := n, M
	for i > 0 && j > 0 {
		if dp[i][j] != dp[i-1][j] {
			res = append(res, i)
			j -= weights[i-1]
		}
		i--
	}

	// Вывод ответа
	fmt.Println(len(res))
	fmt.Println(strings.Trim(fmt.Sprint(res), "[]"))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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
