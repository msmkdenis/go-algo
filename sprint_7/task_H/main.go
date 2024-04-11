package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxFlowers(n, m int, grid [][]int) int {
	// Инициализация массива dp с размерами n x m
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, m)
	}

	// Начальное значение для первой клетки
	dp[0][0] = grid[0][0]

	// Заполняем первую строку
	for j := 1; j < m; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}

	// Заполняем первый столбец
	for i := 1; i < n; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}

	// Заполняем остальные клетки
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			// Выбираем максимальное значение из двух возможных предыдущих путей
			dp[i][j] = max(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}

	// Возвращаем количество цветочков, которое Кондратина сможет собрать
	return dp[n-1][m-1]
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

func main() {
	// Считываем входные данные
	input := getInputData()

	// Парсим размеры поля
	// Парсим размеры поля
	nm := strings.Fields(input[0])
	n, _ := strconv.Atoi(nm[0])
	m, _ := strconv.Atoi(nm[1])

	// Инициализируем сетку для цветов
	grid := make([][]int, n)
	for i := 0; i < n; i++ {
		row := input[i+1]
		grid[i] = make([]int, m)
		for j := 0; j < m; j++ {
			grid[i][j] = int(row[j] - '0')
		}
	}

	// Находим максимальное количество цветочков
	result := maxFlowers(n, m, grid)

	// Выводим результат
	fmt.Println(result)
}
