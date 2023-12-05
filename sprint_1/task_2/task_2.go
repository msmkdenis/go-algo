package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Представьте себе онлайн-игру для поездки в метро: игрок нажимает на кнопку, и на экране появляются три случайных числа.
Если все три числа оказываются одной чётности, игрок выигрывает.

Напишите программу, которая по трём числам определяет, выиграл игрок или нет.
*/

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
	fmt.Println(checkEven(makeIntSlice(s[0])))
}

func makeIntSlice(s string) []int {
	var result []int
	for _, v := range strings.Split(s, " ") {
		i, _ := strconv.Atoi(v)
		result = append(result, i)
	}
	return result
}

func checkEven(nums []int) string {
	var counter int
	for _, v := range nums {
		if v%2 != 0 {
			counter++
		} else {
			counter--
		}
	}

	if counter == len(nums) || -counter == len(nums) {
		return "WIN"
	} else {
		return "FAIL"
	}
}
