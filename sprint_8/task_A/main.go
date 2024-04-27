package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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
	if len(inputData) == 0 {
		fmt.Println("")
		return
	}
	words := strings.Fields(inputData[0])
	left := 0
	right := len(words) - 1
	for left < right {
		words[left], words[right] = words[right], words[left]
		left++
		right--
	}
	fmt.Println(strings.Join(words, " "))
}
