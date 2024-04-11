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

	nm := readArray(input[0])
	n := nm[0]
	m := nm[1]

	flowers := readMatrix(input[1:], n, m)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}

	dp[n-1][0] = flowers[n-1][0]

	for i := n - 1; i >= 0; i-- {
		for j := 0; j < m; j++ {
			dp[i][j] = max(getPrevValueAtIndex(dp, i+1, j), getPrevValueAtIndex(dp, i, j-1)) + flowers[i][j]
		}
	}

	fmt.Println(dp[0][m-1])

	path := ""
	i := 0
	j := m - 1
	for !(j == 0 && i == n-1) {
		if getPrevValueAtIndex(dp, i+1, j) >= getPrevValueAtIndex(dp, i, j-1) && i+1 <= n-1 {
			i++
			path += "U"
		} else {
			j--
			path += "R"
		}
	}

	fmt.Print(Reverse(path))
}

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func getPrevValueAtIndex(dp [][]int, i int, j int) int {
	if i >= len(dp) || j < 0 {
		return 0
	}

	return dp[i][j]
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

func readArray(str string) []int {
	listString := strings.Split(str, " ")
	arr := make([]int, len(listString))
	for i := 0; i < len(listString); i++ {
		arr[i], _ = strconv.Atoi(listString[i])
	}
	return arr
}

func readMatrix(arr []string, rows int, cols int) [][]int {
	matrix := make([][]int, rows)
	for i := 0; i < rows; i++ {
		matrix[i] = make([]int, cols)
		row := arr[i]
		for j, number := range row {
			matrix[i][j] = int(number) - 48
		}
	}
	return matrix
}
