package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := getInputData()
	fmt.Println(allowPassenger(input[0], input[1]))
}

func allowPassenger(baseName, passportName string) string {
	// Создаем хэш-таблицы для подсчета частот символов в обеих строках
	baseHash := make(map[rune]int)
	passportHash := make(map[rune]int)

	// Считаем частоты символов в строках
	for _, char := range baseName {
		baseHash[char]++
	}
	for _, char := range passportName {
		passportHash[char]++
	}

	// Сравниваем частоты символов
	diffCount := 0
	for char, count := range baseHash {
		if count < passportHash[char] {
			count, passportHash[char] = passportHash[char], count
		}

		if count-passportHash[char] > 1 {
			return "FAIL"
		}

		fmt.Println(string(char))
		if count != passportHash[char] {

			diffCount++
		}

		if diffCount > 1 {
			return "FAIL"
		}
	}

	// Если количество отличающихся символов не превышает одного, возвращаем true
	return "OK"
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
