package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	const maxCapacity = 15 * 1024 * 1024
	buf := make([]byte, maxCapacity)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)

	var s []string
	for scanner.Scan() {
		line := scanner.Text()
		s = append(s, line)
		if line == "" {
			break
		}
	}

	p := make(map[int]int)
	p[0] = 0
	capacity, _ := strconv.Atoi(s[0])
	data := makeIntSlice(s[1], capacity)
	if len(data) == 1 {
		fmt.Println(0)
		return
	}

	prefixSum := make([]int, capacity+1)
	prefixSum[0] = 0
	var length int
	for i := 0; i < len(data); i++ {
		elem := 0
		if data[i] == 0 {
			elem = -1
		} else {
			elem = 1
		}
		prefixSum[i+1] = prefixSum[i] + elem
		if _, ok := p[prefixSum[i+1]]; !ok {
			p[prefixSum[i+1]] = i + 1
		} else {
			prev := p[prefixSum[i+1]]
			if length < i-prev {
				length = i - prev + 1
			}
		}
	}

	counter := make([]int, 0, len(prefixSum))
	for i, _ := range prefixSum {
		counter = append(counter, i)
	}

	fmt.Println(length)

}

func makeIntSlice(s string, capacity int) []int {
	result := make([]int, 0, capacity)
	for _, v := range strings.Split(s, " ") {
		i, _ := strconv.Atoi(v)
		result = append(result, i)
	}
	return result
}
