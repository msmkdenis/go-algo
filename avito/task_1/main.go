package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var s []string
	for scanner.Scan() {
		line := scanner.Text()
		s = append(s, line)
		if line == "" {
			break
		}
	}

	qRows, _ := strconv.Atoi(s[0])

	for i := 1; i <= qRows; i++ {
		ss := strings.Split(s[i], " ")
		field := make(map[string]int)
		for _, v := range ss {
			num, ok := field[v]
			if !ok {
				field[v] = 1
			} else {
				field[v] = num + 1
			}
		}
		if check(field) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
		clear(field)
	}
}

func check(field map[string]int) bool {
	for v := range field {
		if v == "4" && field[v] > 1 {
			return false
		}
		if v == "3" && field[v] > 2 {
			return false
		}
		if v == "2" && field[v] > 3 {
			return false
		}
		if v == "1" && field[v] > 4 {
			return false
		}
	}
	return true
}
