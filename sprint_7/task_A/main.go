package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	maxProfit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			maxProfit += prices[i] - prices[i-1]
		}
	}

	return maxProfit
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
	inputData := getInputData()

	n, _ := strconv.Atoi(inputData[0])
	pricesStr := inputData[1]
	prices := make([]int, n)

	pricesSplit := bufio.NewScanner(strings.NewReader(pricesStr))
	pricesSplit.Split(bufio.ScanWords)
	for i := 0; i < n && pricesSplit.Scan(); i++ {
		price, _ := strconv.Atoi(pricesSplit.Text())
		prices[i] = price
	}

	result := maxProfit(prices)
	fmt.Println(result)
}
