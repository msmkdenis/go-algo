package main

import (
	"bufio"
	"fmt"
	"os"
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

	if len(s[0]) != len(s[1]) {
		fmt.Println("NO")
		return
	}

	first := []rune(s[0])
	second := []rune(s[1])
	words := map[rune]rune{}
	wordSet := map[rune]struct{}{}
	for i := 0; i < len(first); i++ {
		if word, ok := words[first[i]]; !ok {
			if _, ok := wordSet[second[i]]; ok {
				fmt.Println("NO")
				return
			}
			wordSet[second[i]] = struct{}{}
			words[first[i]] = second[i]
		} else {
			if word != second[i] {
				fmt.Println("NO")
				return
			}
		}
	}
	fmt.Println("YES")
}
