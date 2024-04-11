package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInputData() (int, int, []int) {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Считываем сумму и количество купюр
	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	// Считываем номиналы купюр
	scanner.Scan()
	line := scanner.Text()
	coinsStr := strings.Fields(line)
	coins := make([]int, n)
	for i, coinStr := range coinsStr {
		coin, _ := strconv.Atoi(coinStr)
		coins[i] = coin
	}

	return m, n, coins
}

func countWays(m int, n int, coins []int) int {
	dp := make([]int, m+1)
	dp[0] = 1

	for _, coin := range coins {
		for i := coin; i <= m; i++ {
			dp[i] += dp[i-coin]
		}
	}

	return dp[m]
}

func main() {
	m, n, coins := getInputData()

	ways := countWays(m, n, coins)
	fmt.Println(ways)
}
