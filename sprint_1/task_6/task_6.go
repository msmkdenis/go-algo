package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
Помогите Васе понять, будет ли фраза палиндромом.
Учитываются только буквы и цифры, заглавные и строчные буквы считаются одинаковыми.

Решение должно работать за O(N), где N — длина строки на входе.
*/

func main() {
	const maxCapacity = 1024 * 1024
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

	sentence := strings.ToLower(strip(s[0]))

	fmt.Println(isPalindrome(sentence))
}

func strip(s string) string {
	var result strings.Builder
	for i := 0; i < len(s); i++ {
		b := s[i]
		if ('a' <= b && b <= 'z') ||
			('A' <= b && b <= 'Z') ||
			('0' <= b && b <= '9') {
			result.WriteByte(b)
		}
	}
	return result.String()
}

func isPalindrome(s string) string {
	left := 0
	right := len(s) - 1
	for left < right {
		if s[left] != s[right] {
			return "False"
		}
		left++
		right--
	}
	return "True"
}
