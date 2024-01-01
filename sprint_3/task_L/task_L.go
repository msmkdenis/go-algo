package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Вася решил накопить денег на два одинаковых велосипеда — себе и сестре.
У Васи есть копилка, в которую каждый день он может добавлять деньги (если, конечно, у него есть такая финансовая возможность).
В процессе накопления Вася не вынимает деньги из копилки.

У вас есть информация о росте Васиных накоплений — сколько у Васи в копилке было денег в каждый из дней.

Ваша задача — по заданной стоимости велосипеда определить

первый день, в которой Вася смог бы купить один велосипед,
и первый день, в который Вася смог бы купить два велосипеда.
Подсказка: решение должно работать за O(log n).
*/

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
	days, _ := strconv.Atoi(s[0])
	savings := makeIntSlice(s[1], days)
	cost, _ := strconv.Atoi(s[2])

	firstDay := binarySearch(savings, cost, 0, days-1)
	secondDay := binarySearch(savings, cost*2, 0, days-1)
	fmt.Println(firstDay, secondDay)
}

func makeIntSlice(s string, cap int) []int {
	result := make([]int, 0, cap)
	for _, v := range strings.Split(s, " ") {
		i, _ := strconv.Atoi(v)
		result = append(result, i)
	}
	return result
}

func binarySearch(arr []int, x int, left int, right int) int {
	fmt.Println(left, right)
	if arr[right] < x {
		return -1
	}

	if left < right {
		mid := (left + right) / 2
		fmt.Println(mid)
		if arr[mid] < x {
			return binarySearch(arr, x, mid+1, right)
		} else {
			return binarySearch(arr, x, left, mid)
		}
	} else {
		return left + 1
	}
}
