package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Считываем входные данные
	input := getInputData()

	// Считываем число слитков и вместимость рюкзака
	line := input[0]
	fields := strings.Fields(line)
	n, _ := strconv.Atoi(fields[0])
	M, _ := strconv.Atoi(fields[1])

	// Считываем массы слитков
	weightsStr := strings.Fields(input[1])
	weights := make([]int, n)
	for i, str := range weightsStr {
		weights[i], _ = strconv.Atoi(str)
	}

	// Инициализируем массив dp для хранения максимальной суммарной массы слитков
	dp := make([]int, M+1)

	// Заполняем массив dp
	for i := 0; i < n; i++ {
		for j := M; j >= weights[i]; j-- {
			dp[j] = max(dp[j], dp[j-weights[i]]+weights[i])
		}
	}

	// Выводим результат
	fmt.Println(dp[M])
}

// Функция для нахождения максимума
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Функция для считывания входных данных из файла
func getInputData() []string {
	inputFile := "input.txt"
	input, _ := os.Open(inputFile)
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
