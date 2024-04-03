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
	stats := strings.Split(input[0], " ")
	v, _ := strconv.Atoi(stats[0])
	edges := make(map[int][]int)
	for i := 1; i < len(input); i++ {
		s1 := strings.Split(input[i], " ")
		v1, _ := strconv.Atoi(s1[0])
		e1, _ := strconv.Atoi(s1[1])
		if _, ok := edges[v1]; !ok {
			edges[v1] = []int{e1}
		} else {
			edges[v1] = append(edges[v1], e1)
		}
	}
	for i := 1; i <= v; i++ {
		if _, ok := edges[i]; !ok {
			fmt.Println(0)
		} else {
			fmt.Printf("%d %s\n", len(edges[i]), printSlice(edges[i]))
		}
	}
}

func printSlice(s []int) string {
	var sb strings.Builder
	for i := 0; i < len(s); i++ {
		if i == len(s)-1 {
			sb.WriteString(strconv.Itoa(s[i]))
		} else {
			sb.WriteString(strconv.Itoa(s[i]) + " ")
		}
	}
	return sb.String()
}

func getInputData() []string {
	input, _ := os.Open("input.txt")
	defer input.Close()

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	var s []string
	for scanner.Scan() {
		bufStr := scanner.Text()
		s = append(s, bufStr)
	}

	return s
}
