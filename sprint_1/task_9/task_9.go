package main

import "fmt"

/*
Напишите программу, которая определяет, будет ли положительное целое число степенью четвёрки.

Подсказка: степенью четвёрки будут все числа вида 4n, где n – целое неотрицательное число.
*/

func main() {
	var number int
	fmt.Scan(&number)
	fmt.Println(isPowerOfFour(number))
}

func isPowerOfFour(num int) string {
	if num == 1 {
		return "True"
	}

	counter := 4
	for counter <= num {
		if counter == num {
			return "True"
		}
		counter *= 4
	}

	return "False"
}
