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

	var length []int
	data := []rune(s[0])
	dict := make(map[rune]struct{})
	first := 0
	second := 0
	for i := 0; i < len(data); i++ {
		if i == 0 {
			dict[data[i]] = struct{}{}
			second++
			continue
		}

		if _, ok := dict[data[i]]; !ok {
			dict[data[i]] = struct{}{}
			second++
		} else {
			for j := first; j < second-first; j++ {
				first++
				if _, ok := dict[data[j]]; ok {
					first++
					delete(dict, data[j])
					break
				}
			}
		}
		length = append(length, second-first)
	}

	fmt.Println(length)

	fmt.Println(second - first)
}
