package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Рита решила оставить у себя одежду только трёх цветов: розового, жёлтого и малинового.
После того как вещи других расцветок были убраны, Рита захотела отсортировать свой новый гардероб по цветам.
Сначала должны идти вещи розового цвета, потом —– жёлтого, и в конце —– малинового.
Помогите Рите справиться с этой задачей.

Примечание: попробуйте решить задачу за один проход по массиву!
*/

func main() {

	const maxCapacity = 5 * 1024 * 1024
	buf := make([]byte, maxCapacity)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)

	var s []string
	for scanner.Scan() {
		line := scanner.Text()
		s = append(s, line)
		if line == "" || line == "0" {
			break
		}
	}

	capacity, _ := strconv.Atoi(s[0])
	if capacity == 0 {
		fmt.Println("")
		return
	}
	wardrobe := makeIntSlice(s[1], capacity)
	sortedWardrobe := countingSort(wardrobe, 3)
	var answer strings.Builder
	for i, v := range sortedWardrobe {
		if i == len(sortedWardrobe)-1 {
			answer.WriteString(strconv.Itoa(v))
		} else {
			answer.WriteString(strconv.Itoa(v) + " ")
		}
	}
	fmt.Println(answer.String())
}

func countingSort(array []int, k int) []int {
	countedValues := make([]int, k)
	for _, value := range array {
		countedValues[value]++
	}

	index := 0
	for value := 0; value < k; value++ {
		for amount := 0; amount < countedValues[value]; amount++ {
			array[index] = value
			index++
		}
	}
	return array
}

func makeIntSlice(s string, cap int) []int {
	result := make([]int, 0, cap)
	for _, v := range strings.Split(s, " ") {
		i, _ := strconv.Atoi(v)
		result = append(result, i)
	}
	return result
}
