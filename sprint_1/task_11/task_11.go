package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

/*
Вася просил Аллу помочь решить задачу. На этот раз по информатике.

Для неотрицательного целого числа X списочная форма –— это массив его цифр слева направо.
К примеру, для 1231 списочная форма будет [1,2,3,1].
На вход подается количество цифр числа Х, списочная форма неотрицательного числа Х и неотрицательное число K.
Число К не превосходят 10000. Длина числа Х не превосходит 1000.

Нужно вернуть списочную форму числа X + K.
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

	formLength, _ := strconv.Atoi(s[0])
	form := makeIntSlice(s[1], formLength)
	number, _ := strconv.Atoi(s[2])
	if formLength < 18 {
		digit := makeDigit(form) + int(number)
		result := makeForm(digit)
		for n, v := range result {
			if n == len(result)-1 {
				fmt.Print(v)
			} else {
				fmt.Print(v, " ")
			}
		}
	} else {
		l := len(strconv.Itoa(number))
		subDigit := form[formLength-l-1:]
		digit := makeDigit(subDigit) + int(number)
		subResult := makeForm(digit)
		form = slices.Delete(form, formLength-l-1, len(form))
		form = append(form, subResult...)
		for n, v := range form {
			if n == len(form)-1 {
				fmt.Print(v)
			} else {
				fmt.Print(v, " ")
			}
		}

	}

}

func makeIntSlice(s string, len int) []int {
	result := make([]int, 0, len)
	for _, v := range strings.Split(s, " ") {
		i, _ := strconv.Atoi(v)
		result = append(result, int(i))
	}
	return result
}

func makeDigit(digits []int) int {
	var number int
	multiplier := 1
	for i := len(digits) - 1; i >= 0; i-- {
		number += digits[i] * multiplier
		multiplier *= 10
	}

	return number
}

func makeForm(number int) []int {
	var result []int
	for number > 0 {
		result = append(result, number%10)
		number /= 10
	}
	slices.Reverse(result)
	return result
}
